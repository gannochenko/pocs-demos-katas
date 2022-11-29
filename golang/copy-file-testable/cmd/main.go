package main

import (
	"copyfiletestable/internal/clients/gcp"
	"copyfiletestable/internal/clients/http"
	"copyfiletestable/internal/services/filecopier"
	"log"
	"os"
)

func main() {
	fileURL := os.Getenv("FILE_URL")
	bucketName := os.Getenv("BUCKET_NAME")
	objectPath := os.Getenv("OBJECT_PATH")

	storageClient := gcp.NewStorageClient()
	httpClient := http.NewClient()

	fileCopier := filecopier.New(&filecopier.ServiceDependencies{
		StorageClient: storageClient,
		HttpClient:    httpClient,
	})

	err := fileCopier.CopyFile(fileURL, bucketName, objectPath)
	if err != nil {
		panic(err)
	}

	log.Println("Done!")
}
