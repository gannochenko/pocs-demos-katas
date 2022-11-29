package http

import (
	"copyfiletestable/internal"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) ExtractFileName(fileURL string) (string, error) {
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		return "", err
	}

	splitPath := strings.Split(parsedURL.Path, "/")

	if len(splitPath) == 0 {
		return "", nil
	}

	return splitPath[len(splitPath)-1], nil
}

func (c *Client) DownloadFile(url string) (handle *internal.DocumentReadHandle, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, errors.New("could not load the file " + fmt.Sprintf("%d", resp.StatusCode))
	}

	return &internal.DocumentReadHandle{
		Reader: resp.Body,
	}, nil
}
