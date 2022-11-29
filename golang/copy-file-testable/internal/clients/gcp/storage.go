package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"copyfiletestable/internal"
	"github.com/pborman/uuid"
	"io"
	"time"
)

type StorageClient struct{}

func NewStorageClient() *StorageClient {
	return &StorageClient{}
}

func (c *StorageClient) UploadFile(handle *internal.DocumentReadHandle, bucketName string, objectPath string) error {
	uploaderCtx := context.Background()

	uploaderCtx, cancel := context.WithTimeout(uploaderCtx, time.Second*50)
	defer cancel()

	client, err := storage.NewClient(uploaderCtx)
	if err != nil {
		return err
	}

	object := client.Bucket(bucketName).Object(objectPath)
	objectWriter := object.NewWriter(uploaderCtx)

	if _, err := io.Copy(objectWriter, handle.Reader); err != nil {
		return err
	}
	if err := objectWriter.Close(); err != nil {
		return err
	}

	return nil
}

func (c *StorageClient) PrepareObjectPath(objectPathPrefix, fileName string) string {
	return objectPathPrefix + "/" + uuid.New() + "-" + fileName
}
