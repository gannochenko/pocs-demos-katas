package http

import (
	"copyfiletestable/internal"
	"github.com/stretchr/testify/mock"
)

type ClientMock struct {
	mock.Mock
}

func NewClientMock() *ClientMock {
	return &ClientMock{}
}

func (m *ClientMock) ExtractFileName(fileURL string) (string, error) {
	args := m.Called(fileURL)
	return args.Get(0).(string), args.Error(1)
}

func (m *ClientMock) DownloadFile(url string) (handle *internal.DocumentReadHandle, err error) {
	args := m.Called(url)

	var arg1 *internal.DocumentReadHandle
	if args.Get(0) != nil {
		arg1 = args.Get(0).(*internal.DocumentReadHandle)
	}

	return arg1, args.Error(1)
}
