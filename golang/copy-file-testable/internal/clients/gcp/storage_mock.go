package gcp

import (
	"copyfiletestable/internal"
	"github.com/stretchr/testify/mock"
)

type StorageClientMock struct {
	mock.Mock
}

func NewStorageClientMock() *StorageClientMock {
	return &StorageClientMock{}
}

func (m *StorageClientMock) UploadFile(handle *internal.DocumentReadHandle, bucketName, objectPath string) error {
	args := m.Called(handle, bucketName, objectPath)
	return args.Error(0)
}

func (m *StorageClientMock) PrepareObjectPath(objectPathPrefix, fileName string) string {
	args := m.Called(objectPathPrefix, fileName)
	return args.Get(0).(string)
}
