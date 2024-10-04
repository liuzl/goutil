package zimage

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	"github.com/nfnt/resize"
)

func AutoCompress(imageData []byte) ([]byte, error) {
	if len(imageData) == 0 {
		return nil, fmt.Errorf("image data is empty")
	}

	// Skip compression for images smaller than 100KB
	if len(imageData) <= 100*1024 {
		return imageData, nil
	}

	// Try to infer the image type
	contentType := InferImageType(imageData)

	// If the image is WebP or an unsupported format, return the original data
	if contentType == "image/webp" || contentType == "application/octet-stream" {
		return imageData, nil
	}

	// Decode the image
	img, format, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	// Get original dimensions
	originalWidth := img.Bounds().Dx()
	originalHeight := img.Bounds().Dy()

	// Resize image if necessary
	var resized image.Image
	if originalWidth > 400 || originalHeight > 400 {
		if originalWidth > originalHeight {
			resized = resize.Resize(400, 0, img, resize.Lanczos3)
		} else {
			resized = resize.Resize(0, 400, img, resize.Lanczos3)
		}
	} else {
		resized = img
	}

	// Encode the image
	var buf bytes.Buffer
	switch format {
	case "jpeg":
		err = jpeg.Encode(&buf, resized, &jpeg.Options{Quality: 80})
	case "png":
		err = png.Encode(&buf, resized)
	case "gif":
		err = gif.Encode(&buf, resized, &gif.Options{})
	default:
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to encode image: %v", err)
	}

	return buf.Bytes(), nil
}
