package image

import (
	"backend/internal/util/syserr"
	"image"
	_ "image/gif"  // Register GIF format
	_ "image/jpeg" // Register JPEG format
	_ "image/png"  // Register PNG format
	"net/http"
)

func DownloadImage(url string) (image.Image, error) {
	// Fetch the image
	resp, err := http.Get(url)
	if err != nil {
		return nil, syserr.Wrap(err, "failed to fetch image")
	}
	defer resp.Body.Close()

	// Check if the HTTP status is OK
	if resp.StatusCode != http.StatusOK {
		return nil, syserr.NewInternal("received non-200 status code", syserr.F("code", resp.StatusCode))
	}

	// Decode the image
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, syserr.Wrap(err, "failed to decode image")
	}

	return img, nil
}
