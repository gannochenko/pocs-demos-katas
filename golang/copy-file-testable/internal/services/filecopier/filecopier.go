package filecopier

import "copyfiletestable/internal"

type gcpStorageClient interface {
	UploadFile(handle *internal.DocumentReadHandle, bucketName string, objectPath string) error
	PrepareObjectPath(objectPathPrefix, fileName string) string
}

type httpClient interface {
	ExtractFileName(fileURL string) (string, error)
	DownloadFile(url string) (handle *internal.DocumentReadHandle, err error)
}

type ServiceDependencies struct {
	StorageClient gcpStorageClient
	HttpClient    httpClient
}

type Service struct {
	storageClient gcpStorageClient
	httpClient    httpClient
}

func New(deps *ServiceDependencies) *Service {
	return &Service{
		storageClient: deps.StorageClient,
		httpClient:    deps.HttpClient,
	}
}

func (s *Service) CopyFile(fileURL, bucketName, objectPath string) (err error) {
	fileName, err := s.httpClient.ExtractFileName(fileURL)
	if err != nil {
		return err
	}

	documentHandle, err := s.httpClient.DownloadFile(fileURL)
	if err != nil {
		return err
	}
	defer func() {
		err := documentHandle.Reader.Close()
		if err != nil {
		}
	}()

	targetObjectPath := s.storageClient.PrepareObjectPath(objectPath, fileName)
	err = s.storageClient.UploadFile(documentHandle, bucketName, targetObjectPath)
	if err != nil {
		return err
	}

	return
}
