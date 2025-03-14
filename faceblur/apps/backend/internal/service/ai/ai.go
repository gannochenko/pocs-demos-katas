package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
	ort "github.com/yalue/onnxruntime_go"
)

// Detection holds information about detected faces
type Detection struct {
	Box       [4]float32 // [x1, y1, x2, y2]
	Confidence float32
}

// FaceDetector encapsulates the YOLO model and session
type FaceDetector struct {
	session      *ort.AdvancedSession
	inputTensor  ort.Value
	outputTensor ort.Value
	inputName    string
	outputName   string
	inputWidth   int
	inputHeight  int
	inputShape   *ort.Shape
}

// NewFaceDetector creates and initializes a new face detector
func NewFaceDetector(modelPath string, width, height int) (*FaceDetector, error) {
	// Initialize ONNX Runtime (should only be done once in your application)
	err := ort.InitializeEnvironment()
	if err != nil {
		return nil, fmt.Errorf("error initializing environment: %v", err)
	}

	// Create a new detector
	detector := &FaceDetector{
		inputWidth:  width,
		inputHeight: height,
	}

	// Define input shape: [batch_size, channels, height, width] (NCHW format)
	sh := ort.NewShape(1, 3, int64(height), int64(width))
	detector.inputShape = &sh

	// Create input tensor (will be reused)
	inputTensor, err := ort.NewEmptyTensor[float32](*detector.inputShape)
	if err != nil {
		ort.DestroyEnvironment()
		return nil, fmt.Errorf("error creating input tensor: %v", err)
	}
	detector.inputTensor = inputTensor

	// Create session options
	sessionOptions, err := ort.NewSessionOptions()
	if err != nil {
		inputTensor.Destroy()
		ort.DestroyEnvironment()
		return nil, fmt.Errorf("error creating session options: %v", err)
	}
	defer sessionOptions.Destroy()

	// Get model metadata to determine input/output names
	metadata, err := ort.GetModelMetadata(modelPath)
	if err != nil {
		inputTensor.Destroy()
		ort.DestroyEnvironment()
		return nil, fmt.Errorf("error getting model metadata: %v", err)
	}

	if len(metadata.Inputs) == 0 || len(metadata.Outputs) == 0 {
		inputTensor.Destroy()
		ort.DestroyEnvironment()
		return nil, fmt.Errorf("model has no inputs or outputs according to metadata")
	}

	detector.inputName = metadata.Inputs[0].Name
	detector.outputName = metadata.Outputs[0].Name

	fmt.Printf("Model input name: %s, output name: %s\n", detector.inputName, detector.outputName)

	// Create output tensor (will be reused)
	outputTensor, err := ort.NewEmptyTensor[float32](ort.NewShape(1, 100, 85))
	if err != nil {
		inputTensor.Destroy()
		ort.DestroyEnvironment()
		return nil, fmt.Errorf("error creating output tensor: %v", err)
	}
	detector.outputTensor = outputTensor

	// Create a session (this is where the model is loaded)
	session, err := ort.NewAdvancedSession(
		modelPath,
		[]string{detector.inputName},
		[]string{detector.outputName},
		[]ort.Value{detector.inputTensor},
		[]ort.Value{detector.outputTensor},
		sessionOptions,
	)
	if err != nil {
		inputTensor.Destroy()
		outputTensor.Destroy()
		ort.DestroyEnvironment()
		return nil, fmt.Errorf("error creating session: %v", err)
	}
	detector.session = session

	return detector, nil
}

// Detect processes an image and returns face detections
func (fd *FaceDetector) Detect(imagePath string, confidenceThreshold float32) ([]Detection, image.Image, error) {
	// Load and preprocess image
	prepImg, originalImage, err := fd.loadAndPreprocessImage(imagePath)
	if err != nil {
		return nil, nil, fmt.Errorf("error loading image: %v", err)
	}

	// Copy data to input tensor
	inputData := fd.inputTensor.GetData().([]float32)
	copy(inputData, prepImg.Data)

	// Run inference
	err = fd.session.Run()
	if err != nil {
		return nil, originalImage, fmt.Errorf("error running inference: %v", err)
	}

	// Process detections
	detections := fd.processOutput(prepImg.OriginalWidth, prepImg.OriginalHeight, confidenceThreshold)

	return detections, originalImage, nil
}

// Close releases resources used by the detector
func (fd *FaceDetector) Close() {
	if fd.session != nil {
		fd.session.Destroy()
	}
	if fd.inputTensor != nil {
		fd.inputTensor.Destroy()
	}
	if fd.outputTensor != nil {
		fd.outputTensor.Destroy()
	}
	// Note: Environment should only be destroyed when application exits
	// ort.DestroyEnvironment()
}

// PreprocessedImage contains preprocessed image data
type PreprocessedImage struct {
	Data           []float32
	OriginalWidth  int
	OriginalHeight int
}

// loadAndPreprocessImage loads an image and preprocesses it for the YOLO model
func (fd *FaceDetector) loadAndPreprocessImage(path string) (*PreprocessedImage, image.Image, error) {
	// Open the image file
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	// Decode the image
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, nil, err
	}

	// Get original dimensions
	origWidth := img.Bounds().Max.X
	origHeight := img.Bounds().Max.Y

	// Resize to expected input size
	resizedImg := resize.Resize(uint(fd.inputWidth), uint(fd.inputHeight), img, resize.Lanczos3)

	// Create NCHW tensor array (1, 3, height, width)
	dataSize := 1 * 3 * fd.inputHeight * fd.inputWidth
	inputData := make([]float32, dataSize)

	// Convert image to RGB and normalize to [0,1]
	for y := 0; y < fd.inputHeight; y++ {
		for x := 0; x < fd.inputWidth; x++ {
			r, g, b, _ := resizedImg.At(x, y).RGBA()
			// Map from [0,65535] to [0,1]
			inputData[0*fd.inputHeight*fd.inputWidth+y*fd.inputWidth+x] = float32(r) / 65535.0
			inputData[1*fd.inputHeight*fd.inputWidth+y*fd.inputWidth+x] = float32(g) / 65535.0
			inputData[2*fd.inputHeight*fd.inputWidth+y*fd.inputWidth+x] = float32(b) / 65535.0
		}
	}

	return &PreprocessedImage{
		Data:           inputData,
		OriginalWidth:  origWidth,
		OriginalHeight: origHeight,
	}, img, nil
}

// processOutput converts model output to detection coordinates
func (fd *FaceDetector) processOutput(origWidth, origHeight int, confidenceThreshold float32) []Detection {
	// Get output data
	outputData := fd.outputTensor.GetData().([]float32)
	shape := fd.outputTensor.GetShape()

	var detections []Detection

	// Process based on output shape
	if len(shape) == 3 {
		// Format: [1, N, 5+C]
		numDetections := int(shape[1])
		valuesPerDetection := int(shape[2])
		
		for i := 0; i < numDetections; i++ {
			baseIdx := i * valuesPerDetection
			
			// Confidence score location depends on YOLO version
			confidence := outputData[baseIdx+4]
			
			if confidence > confidenceThreshold {
				// YOLO outputs normalized [x_center, y_center, width, height]
				x := outputData[baseIdx+0]
				y := outputData[baseIdx+1]
				w := outputData[baseIdx+2]
				h := outputData[baseIdx+3]
				
				// Convert to absolute coordinates (un-normalize)
				x1 := (x - w/2) * float32(origWidth)
				y1 := (y - h/2) * float32(origHeight)
				x2 := (x + w/2) * float32(origWidth)
				y2 := (y + h/2) * float32(origHeight)
				
				detections = append(detections, Detection{
					Box:        [4]float32{x1, y1, x2, y2},
					Confidence: confidence,
				})
			}
		}
	} else if len(shape) == 2 {
		// Format: [N, 5+C] (flattened output)
		numDetections := int(shape[0])
		valuesPerDetection := int(shape[1])
		
		for i := 0; i < numDetections; i++ {
			baseIdx := i * valuesPerDetection
			
			confidence := outputData[baseIdx+4]
			
			if confidence > confidenceThreshold {
				x := outputData[baseIdx+0]
				y := outputData[baseIdx+1]
				w := outputData[baseIdx+2]
				h := outputData[baseIdx+3]
				
				x1 := (x - w/2) * float32(origWidth)
				y1 := (y - h/2) * float32(origHeight)
				x2 := (x + w/2) * float32(origWidth)
				y2 := (y + h/2) * float32(origHeight)
				
				detections = append(detections, Detection{
					Box:        [4]float32{x1, y1, x2, y2},
					Confidence: confidence,
				})
			}
		}
	} else {
		fmt.Printf("Unsupported output shape: %v\n", shape)
	}

	return detections
}

// drawBoundingBox draws a bounding box on the image
func drawBoundingBox(img image.Image, detection Detection) {
	// Convert to drawable image
	bounds := img.Bounds()
	rgbaImg, ok := img.(*image.RGBA)
	if !ok {
		// Convert image to RGBA
		rgbaImg = image.NewRGBA(bounds)
		draw.Draw(rgbaImg, bounds, img, bounds.Min, draw.Src)
	}

	// Get bounding box coordinates
	x1 := int(detection.Box[0])
	y1 := int(detection.Box[1])
	x2 := int(detection.Box[2])
	y2 := int(detection.Box[3])

	// Ensure coordinates are within image bounds
	if x1 < 0 {
		x1 = 0
	}
	if y1 < 0 {
		y1 = 0
	}
	if x2 >= bounds.Max.X {
		x2 = bounds.Max.X - 1
	}
	if y2 >= bounds.Max.Y {
		y2 = bounds.Max.Y - 1
	}

	// Draw bounding box (red)
	c := color.RGBA{255, 0, 0, 255}
	
	// Draw horizontal lines
	for x := x1; x <= x2; x++ {
		rgbaImg.Set(x, y1, c)
		rgbaImg.Set(x, y2, c)
	}
	
	// Draw vertical lines
	for y := y1; y <= y2; y++ {
		rgbaImg.Set(x1, y, c)
		rgbaImg.Set(x2, y, c)
	}
}

// saveImage saves the image to a file
func saveImage(img image.Image, outputPath string) error {
	// Create output file
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Encode and save as JPEG
	return jpeg.Encode(f, img, &jpeg.Options{Quality: 90})
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <path-to-onnx-model> <path-to-image>")
		return
	}

	modelPath := os.Args[1]
	imagePath := os.Args[2]

	// Create face detector (model is loaded here only once)
	detector, err := NewFaceDetector(modelPath, 416, 416) // Adjust dimensions as needed
	if err != nil {
		log.Fatalf("Error initializing face detector: %v", err)
	}
	defer detector.Close() // Clean up resources when done

	// Detect faces
	detections, originalImage, err := detector.Detect(imagePath, 0.5)
	if err != nil {
		log.Fatalf("Error detecting faces: %v", err)
	}

	fmt.Printf("Detected %d faces\n", len(detections))
	for i, det := range detections {
		fmt.Printf("Face #%d: Box=[%.1f, %.1f, %.1f, %.1f], Confidence=%.2f\n",
			i+1, det.Box[0], det.Box[1], det.Box[2], det.Box[3], det.Confidence)
		
		// Draw bounding box on the original image
		drawBoundingBox(originalImage, det)
	}

	// Save the result
	outputPath := "output.jpg"
	saveImage(originalImage, outputPath)
	fmt.Printf("Result saved to %s\n", outputPath)

	// If you have multiple images to process, you would just call detector.Detect() 
	// multiple times with different image paths instead of creating a new detector each time
}