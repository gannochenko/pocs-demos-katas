package main

import (
    "io"
    "net/http"
    "net/url"
    "errors"
    "os"
    "fmt"
    "log"
    "time"
    "context"
    "strings"

     "github.com/pborman/uuid"
     "cloud.google.com/go/storage"
)

func main() {
    fileURL := os.Getenv("FILE_URL")
    bucketName := os.Getenv("BUCKET_NAME")
    objectPath := os.Getenv("OBJECT_PATH")

    fileName, err := extractFileName(fileURL)
    if err != nil {
        panic(err)
    }

    fileReader, closeReader, err := downloadFile(fileURL)
    if err != nil {
        panic(err)
    }
    defer func() {
       err := closeReader()
       if err != nil {
           panic(err)
       }
    }()

    uploadFile(fileName, fileReader, bucketName, objectPath)

    log.Println("Done!")
}

func downloadFile(url string) (reader io.ReadCloser, Close func() error, err error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, nil, err
    }

    client := &http.Client{
        Transport: &http.Transport{},
    }

    resp, err := client.Do(req)
    if err != nil {
        return nil, nil, err
    }

    if resp.StatusCode != http.StatusOK {
        return nil, nil, errors.New("could not load the file "+fmt.Sprintf("%s", resp.StatusCode))
    }

    return resp.Body, resp.Body.Close, nil
}

func uploadFile(fileName string, fileReader io.ReadCloser, bucketName string, objectPath string) error {
    uploaderCtx := context.Background()

    uploaderCtx, cancel := context.WithTimeout(uploaderCtx, time.Second*50)
    defer cancel()

    targetObjectPath := objectPath+"/"+uuid.New()+"-"+fileName
    log.Println("Uploading to "+bucketName+"/"+targetObjectPath)

    client, err := storage.NewClient(uploaderCtx)
    if err != nil {
        return err
    }

    object := client.Bucket(bucketName).Object(targetObjectPath)
    objectWriter := object.NewWriter(uploaderCtx)

    if _, err := io.Copy(objectWriter, fileReader); err != nil {
       return err
    }
    if err := objectWriter.Close(); err != nil {
       return err
    }

    return nil
}

func extractFileName(fileURL string) (string, error) {
    parsedURL, err := url.Parse(fileURL)
    if err != nil {
        return "", err
    }

    splitPath := strings.Split(parsedURL.Path, "/")

    if len(splitPath) == 0 {
        return "", nil
    }

    return splitPath[len(splitPath) - 1], nil
}
