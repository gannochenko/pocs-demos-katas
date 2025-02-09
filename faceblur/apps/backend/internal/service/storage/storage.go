package storage

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"encoding/json"

	"github.com/samber/lo"

	"backend/interfaces"
	"backend/internal/util/syserr"

	"cloud.google.com/go/storage"
)

type Service struct {
	configService interfaces.ConfigService
	loggerService interfaces.LoggerService
}

func NewStorageService(configService interfaces.ConfigService, loggerService interfaces.LoggerService) *Service {
	return &Service{
		configService: configService,
		loggerService: loggerService,
	}
}

func (s *Service) GetWriter(ctx context.Context, bucketName string, objectName string) (io.WriteCloser, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, syserr.Wrap(err, "could not create storage client")
	}

	return client.Bucket(bucketName).Object(objectName).NewWriter(ctx), nil
}

func (s *Service) PrepareSignedURL(ctx context.Context, bucketName string, objectName string, ttl time.Duration, method string, contentType string) (url string, err error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", syserr.Wrap(err, "could not create storage client")
	}
	defer func() {
		localErr := client.Close()
		if localErr != nil {
			s.loggerService.LogError(ctx, localErr)
		}
	}()

	config, err := s.configService.GetConfig()
	if err != nil {
		return "", syserr.Wrap(err, "could not load config")
	}

	clientEmail := ""
	var privateKey []byte
	var signBytes func(bytes []byte) ([]byte, error)
	insecure := os.Getenv("STORAGE_EMULATOR_HOST") != ""
	if insecure {
		clientEmail = "foo@bar.baz"
		signBytes = func(bytes []byte) ([]byte, error) {
			return bytes, nil
		}
	} else {
		clientEmail, privateKey, err = s.extractCredentials(config.GCP.ServiceAccount)
		if err != nil {
			return "", syserr.Wrap(err, "could not extract credentials")
		}
		signBytes = nil
	}

	if method == "" {
		method = http.MethodGet
	}

	return client.Bucket(bucketName).SignedURL(objectName, &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		Method:         method,
		Expires:        time.Now().Add(ttl),
		GoogleAccessID: clientEmail,
		PrivateKey:     privateKey,
		Insecure:       insecure,
		ContentType:    contentType,
		SignBytes:      signBytes,
	})
}

func (s *Service) GetPublicURL(bucketName string, objectName string) string {
	emulatorHost := os.Getenv("STORAGE_EMULATOR_HOST")
	host := lo.Ternary(emulatorHost != "", emulatorHost, "https://storage.googleapis.com")

	return fmt.Sprintf("%s/%s/%s", host, bucketName, url.PathEscape(objectName))
}

func (s *Service) extractCredentials(serviceAccount string) (clientEmail string, privateKey []byte, err error) {
	var credentials struct {
		ClientEmail string `json:"client_email"`
		PrivateKey  string `json:"private_key"`
	}

	err = json.Unmarshal([]byte(serviceAccount), &credentials)
	if err != nil {
		return "", nil, err
	}

	return credentials.ClientEmail, []byte(credentials.PrivateKey), nil
}
