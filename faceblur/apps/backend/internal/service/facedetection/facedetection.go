package facedetection

import (
	"backend/interfaces"
	"backend/internal/domain"
	imageUtil "backend/internal/util/image"
	"backend/internal/util/logger"
	"backend/internal/util/syserr"
	"context"
	"fmt"
	"image"
	"runtime"
	"sync"

	"github.com/nfnt/resize"
	ort "github.com/yalue/onnxruntime_go"
)

type modelSession struct {
	Session *ort.AdvancedSession
	Input   *ort.Tensor[float32]
	Output  *ort.Tensor[float32]
}

func (m *modelSession) Destroy() {
	m.Session.Destroy()
	m.Input.Destroy()
	m.Output.Destroy()
}

type Service struct {
	configService interfaces.ConfigService
	loggerService interfaces.LoggerService

	ortCreated bool
	ortCreationMutex sync.Mutex
}

func NewService(configService interfaces.ConfigService, loggerService interfaces.LoggerService) *Service {
	return &Service{
		configService: configService,
		loggerService: loggerService,
		ortCreationMutex: sync.Mutex{},
		ortCreated: false,
	}
}

func (s *Service) Detect(ctx context.Context, image image.Image) ([]*domain.BoundingBox, error) {
	originalWidth := image.Bounds().Canon().Dx()
	originalHeight := image.Bounds().Canon().Dy()

	modelSession, e := s.initSession()
	if e != nil {
		return nil, syserr.Wrap(e, "could not create model session")
	}
	defer modelSession.Destroy()

	e = s.prepareInput(image, modelSession.Input)
	if e != nil {
		return nil, syserr.Wrap(e, "could not convert image to network input")
	}

	e = modelSession.Session.Run()
	if e != nil {
		return nil, syserr.Wrap(e, "could not run session")
	}

	boundingBoxes := s.processOutput(ctx, modelSession.Output.GetData(), originalWidth, originalHeight)
	boundingBoxes = imageUtil.FilterDistinctBoxes(boundingBoxes, 0.45, 0.9, 0)

	// for i, boundingBox := range boundingBoxes {
	// 	fmt.Printf("Box %d: %s\n", i, &boundingBox)
	// }

	result := make([]*domain.BoundingBox, len(boundingBoxes))
	for i, boundingBox := range boundingBoxes {
		result[i] = &domain.BoundingBox{
			X1: boundingBox.X1,
			Y1: boundingBox.Y1,
			X2: boundingBox.X2,
			Y2: boundingBox.Y2,
		}
	}

	return result, nil
}

func (s *Service) prepareInput(pic image.Image, dst *ort.Tensor[float32]) error {
	data := dst.GetData()
	channelSize := 640 * 640
	if len(data) < (channelSize * 3) {
		return syserr.NewInternal(
			fmt.Sprint("Destination tensor only holds %d floats, needs %d (make sure it's the right shape!)", len(data), channelSize*3)		)
	}
	redChannel := data[0:channelSize]
	greenChannel := data[channelSize : channelSize*2]
	blueChannel := data[channelSize*2 : channelSize*3]

	// Resize the image to 640x640 using Lanczos3 algorithm
	pic = resize.Resize(640, 640, pic, resize.Lanczos3)
	i := 0
	for y := 0; y < 640; y++ {
		for x := 0; x < 640; x++ {
			r, g, b, _ := pic.At(x, y).RGBA()
			redChannel[i] = float32(r>>8) / 255.0
			greenChannel[i] = float32(g>>8) / 255.0
			blueChannel[i] = float32(b>>8) / 255.0
			i++
		}
	}

	return nil
}

func (s *Service) getSharedLibPath() (string, error) {
	if runtime.GOOS == "windows" {
		if runtime.GOARCH == "amd64" {
			return "./third_party/onnxruntime.dll", nil
		}
	}
	if runtime.GOOS == "darwin" {
		if runtime.GOARCH == "arm64" {
			return "./third_party/onnxruntime_arm64.dylib", nil
		}
		if runtime.GOARCH == "amd64" {
			return "./third_party/onnxruntime_amd64.dylib", nil
		}
	}
	if runtime.GOOS == "linux" {
		if runtime.GOARCH == "arm64" {
			return "./third_party/onnxruntime_arm64.so", nil
		}
		return "./third_party/onnxruntime.so", nil
	}

	return "", syserr.NewInternal("unable to find a version of the onnxruntime library supporting this system")
}

func (s *Service) processOutput(ctx context.Context, output []float32, originalWidth, originalHeight int) []imageUtil.BoundingBox {
	boundingBoxes := make([]imageUtil.BoundingBox, 0)

	detectionCount := len(output) / 5 // Ensure we don't go out of bounds

	for idx := 0; idx < detectionCount; idx++ {
		if idx >= len(output) || 8400+idx >= len(output) || 2*8400+idx >= len(output) || 3*8400+idx >= len(output) {
			s.loggerService.Info(ctx, "skipping idx out of bounds", logger.F("idx", idx))
			continue
		}

		xc, yc := output[idx], output[8400+idx]
		w, h := output[2*8400+idx], output[3*8400+idx]
		confidence := output[4*8400+idx]

		if confidence < 0.5 {
			continue
		}

		x1 := (xc - w/2) / 640 * float32(originalWidth)
		y1 := (yc - h/2) / 640 * float32(originalHeight)
		x2 := (xc + w/2) / 640 * float32(originalWidth)
		y2 := (yc + h/2) / 640 * float32(originalHeight)

		boundingBoxes = append(boundingBoxes, imageUtil.BoundingBox{
			Confidence: confidence,
			X1:         x1,
			Y1:         y1,
			X2:         x2,
			Y2:         y2,
		})
	}

	return boundingBoxes
}

func (s *Service) initORT() error {
	if s.ortCreated {
		return nil
	}

	s.ortCreationMutex.Lock()
	defer s.ortCreationMutex.Unlock()

	libraryPath, err := s.getSharedLibPath()
	if err != nil {
		return syserr.Wrap(err, "could not get library path")
	}

	ort.SetSharedLibraryPath(libraryPath)
	err = ort.InitializeEnvironment()
	if err != nil {
		return syserr.Wrap(err, "error initializing ORT environment")
	}

	s.ortCreated = true

	return nil
}

func (s *Service) initSession() (*modelSession, error) {
	config, err := s.configService.GetConfig()
	if err != nil {
		return nil, syserr.Wrap(err, "could not get config")
	}

	err = s.initORT()
	if err != nil {
		return nil, syserr.Wrap(err, "could not initialize ORT")
	}

	inputShape := ort.NewShape(1, 3, 640, 640)
	inputTensor, err := ort.NewEmptyTensor[float32](inputShape)
	if err != nil {
		return nil, syserr.Wrap(err, "error creating input tensor")
	}
	outputShape := ort.NewShape(1, 5, 8400)
	outputTensor, err := ort.NewEmptyTensor[float32](outputShape)
	if err != nil {
		inputTensor.Destroy()
		return nil, syserr.Wrap(err, "error creating output tensor")
	}
	options, err := ort.NewSessionOptions()
	if err != nil {
		inputTensor.Destroy()
		outputTensor.Destroy()
		return nil, syserr.Wrap(err, "error creating ORT session options")
	}
	defer options.Destroy()

	// If CoreML is enabled, append the CoreML execution provider
	if config.Backend.Worker.UseCoreML {
		err = options.AppendExecutionProviderCoreML(0)
		if err != nil {
			inputTensor.Destroy()
			outputTensor.Destroy()
			return nil, syserr.Wrap(err, "error enabling CoreML")
		}
	}

	session, err := ort.NewAdvancedSession(config.Backend.Worker.ModelPath,
		[]string{"images"}, []string{"output0"},
		[]ort.ArbitraryTensor{inputTensor},
		[]ort.ArbitraryTensor{outputTensor},
		options)
	if err != nil {
		inputTensor.Destroy()
		outputTensor.Destroy()
		return nil, syserr.Wrap(err, "error creating ORT session")
	}

	return &modelSession{
		Session: session,
		Input:   inputTensor,
		Output:  outputTensor,
	}, nil
}
