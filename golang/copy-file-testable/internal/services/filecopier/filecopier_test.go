package filecopier_test

import (
	"io"
	"testing"

	"copyfiletestable/internal"
	"copyfiletestable/internal/clients/gcp"
	"copyfiletestable/internal/clients/http"
	ioMock "copyfiletestable/internal/io"
	"copyfiletestable/internal/services/filecopier"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCopyFile(t *testing.T) {
	createMocks := func() (*gcp.StorageClientMock, *ioMock.ReaderMock, *http.ClientMock) {
		storageMock := gcp.NewStorageClientMock()
		storageMock.On("UploadFile", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)
		storageMock.On("PrepareObjectPath", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("")

		readerMock := ioMock.NewReaderMock()
		readerMock.On("Read", mock.Anything).Return(0, io.EOF)
		readerMock.On("Close").Return(nil)

		httpClient := http.NewClientMock()
		httpClient.On("ExtractFileName", mock.AnythingOfType("string")).Return("", nil)
		httpClient.On("DownloadFile", mock.AnythingOfType("string")).Return(&internal.DocumentReadHandle{
			Reader: readerMock,
		}, nil)

		return storageMock, readerMock, httpClient
	}

	type setup struct {
		FileURL     string
		BucketName  string
		ObjectName  string
		StorageMock *gcp.StorageClientMock
		ReaderMock  *ioMock.ReaderMock
		ClientMock  *http.ClientMock
	}

	type setupFunc func(t *testing.T) *setup
	type verifyFunc func(t *testing.T, setup *setup, err error)

	testCases := map[string]struct {
		setupFunc  setupFunc
		verifyFunc verifyFunc
	}{
		"Should return, unpaginated": {
			setupFunc: func(t *testing.T) *setup {
				storageMock, readerMock, httpClient := createMocks()
				return &setup{
					FileURL:     "http://path.to/file.pdf",
					BucketName:  "my-bucket",
					ObjectName:  "my-object",
					StorageMock: storageMock,
					ReaderMock:  readerMock,
					ClientMock:  httpClient,
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, err error) {
				assert.NoError(t, err)
				setup.ClientMock.AssertCalled(t, "DownloadFile", mock.Anything)
				setup.StorageMock.AssertCalled(t, "UploadFile", mock.Anything, mock.Anything, mock.Anything)
				setup.ReaderMock.AssertCalled(t, "Close")
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			setup := tc.setupFunc(t)

			fileCopier := filecopier.New(&filecopier.ServiceDependencies{
				HttpClient:    setup.ClientMock,
				StorageClient: setup.StorageMock,
			})

			err := fileCopier.CopyFile(setup.FileURL, setup.BucketName, setup.ObjectName)

			tc.verifyFunc(t, setup, err)
		})
	}
}
