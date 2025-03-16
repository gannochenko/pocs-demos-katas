package image

import (
	"backend/internal/domain"
	"backend/internal/util/syserr"
	"bytes"
	"image"
	_ "image/gif" // Register GIF format
	"image/jpeg"
	_ "image/jpeg" // Register JPEG format
	"image/png"
	_ "image/png" // Register PNG format
	"net/http"
	"strings"

	"github.com/disintegration/imaging"
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

func BlurBoxes(img image.Image, boxes []*domain.BoundingBox, blurRadius float64) (image.Image, error) {
	// Create a clone of the original image to modify
	dst := imaging.Clone(img)

	// Process each bounding box
	for _, box := range boxes {
		// Convert float32 coordinates to integers
		x1, y1 := int(box.X1), int(box.Y1)
		x2, y2 := int(box.X2), int(box.Y2)
		
		// Ensure coordinates are within bounds
		bounds := dst.Bounds()
		x1 = max(x1, bounds.Min.X)
		y1 = max(y1, bounds.Min.Y)
		x2 = min(x2, bounds.Max.X)
		y2 = min(y2, bounds.Max.Y)
		
		// Skip invalid boxes
		if x2 <= x1 || y2 <= y1 {
			continue
		}
		
		// Extract the region to blur
		region := imaging.Crop(dst, image.Rect(x1, y1, x2, y2))
		
		// Apply Gaussian blur to the region
		blurred := imaging.Blur(region, blurRadius)
		
		// Paste the blurred region back
		dst = imaging.Paste(dst, blurred, image.Pt(x1, y1))
	}

	return dst, nil
}

func EncodeImage(img image.Image, format string, quality int) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	var err error

	format = strings.ToLower(format)
	switch format {
	case "jpeg", "jpg":
		err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality})
	case "png":
		err = png.Encode(&buf, img)
	default:
		return nil, syserr.NewInternal("unsupported format", syserr.F("format", format))
	}

	if err != nil {
		return nil, syserr.Wrap(err, "error encoding image")
	}

	return &buf, nil
}
