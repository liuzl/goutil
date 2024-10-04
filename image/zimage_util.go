package zimage

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/carlmjohnson/requests"
)

func FetchImage(url string) ([]byte, string, error) {
	var buf bytes.Buffer
	var conntentType string
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	err := requests.URL(url).ToBytesBuffer(&buf).
		AddValidator(func(r *http.Response) error {
			conntentType = r.Header.Get("Content-Type")
			if !strings.HasPrefix(conntentType, "image/") {
				return fmt.Errorf("invalid content type: %s", conntentType)
			}
			return nil
		}).Fetch(ctx)
	if err != nil {
		if err == context.DeadlineExceeded {
			return nil, conntentType, fmt.Errorf("timeout while downloading the image")
		}
		return nil, conntentType, fmt.Errorf("failed to download image: %w", err)
	}

	if buf.Len() == 0 {
		return nil, conntentType, fmt.Errorf("downloaded image is empty")
	}

	return buf.Bytes(), conntentType, nil
}

func ConvertToInlineImageData(imageData []byte, contentType string) string {
	if contentType == "" {
		contentType = InferImageType(imageData)
	}
	base64Data := base64.StdEncoding.EncodeToString(imageData)
	return fmt.Sprintf("data:%s;base64,%s", contentType, base64Data)
}

func InferImageType(data []byte) string {
	if len(data) < 8 {
		return "application/octet-stream"
	}

	switch {
	case bytes.HasPrefix(data, []byte("\xFF\xD8\xFF")):
		return "image/jpeg"
	case bytes.HasPrefix(data, []byte("\x89PNG\r\n\x1a\n")):
		return "image/png"
	case bytes.HasPrefix(data, []byte("GIF87a")) || bytes.HasPrefix(data, []byte("GIF89a")):
		return "image/gif"
	case bytes.HasPrefix(data, []byte("RIFF")) && bytes.Equal(data[8:12], []byte("WEBP")):
		return "image/webp"
	default:
		return "application/octet-stream"
	}
}

func URLToInlineImageData(url string) (string, error) {
	imageData, contentType, err := FetchImage(url)
	if err != nil {
		return "", err
	}
	return ConvertToInlineImageData(imageData, contentType), nil
}
