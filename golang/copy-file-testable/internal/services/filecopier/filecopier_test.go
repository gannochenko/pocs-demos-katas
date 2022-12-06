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

type testFlow func(t *testing.T)

func TestCopyFile(t *testing.T) {
	testCases := map[string]struct {
		testFlow testFlow
	}{
		"Should download a file": {
			testFlow: func(t *testing.T) {
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

				fileCopier := filecopier.New(&filecopier.ServiceDependencies{
					HttpClient:    httpClient,
					StorageClient: storageMock,
				})

				err := fileCopier.CopyFile("http://path.to/file.pdf", "my-bucket", "my-object")

				assert.NoError(t, err)
				httpClient.AssertCalled(t, "DownloadFile", mock.Anything)
				storageMock.AssertCalled(t, "UploadFile", mock.Anything, mock.Anything, mock.Anything)
				readerMock.AssertCalled(t, "Close")
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			tc.testFlow(t)
		})
	}
}
