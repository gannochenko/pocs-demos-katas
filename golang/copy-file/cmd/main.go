package main

import (
    "io"
    "net/http"
    "errors"
    "os"
    "fmt"
    "log"

//     "github.com/pborman/uuid"
//     "cloud.google.com/go/storage"
)

func main() {
    fileURL := os.Getenv("FILE_URL")

    _, closeReader, err := downloadFile(fileURL)
    if err != nil {
        os.Exit(1)
    }
    defer func() {
       err := closeReader()
       if err != nil {
           os.Exit(1)
       }
    }()

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
