package zimage

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchImage(t *testing.T) {
	// 创建一个模拟的 HTTP 服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/jpeg")
		w.Write([]byte("fake image data"))
	}))
	defer server.Close()

	// 测试成功的情况
	imageData, contentType, err := FetchImage(server.URL)
	if err != nil {
		t.Errorf("FetchImage failed: %v", err)
	}
	if contentType != "image/jpeg" {
		t.Errorf("Expected content type image/jpeg, got %s", contentType)
	}
	if !bytes.Equal(imageData, []byte("fake image data")) {
		t.Errorf("Unexpected image data")
	}

	// 测试无效的内容类型
	invalidServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("not an image"))
	}))
	defer invalidServer.Close()

	_, _, err = FetchImage(invalidServer.URL)
	if err == nil {
		t.Error("Expected error for invalid content type, got nil")
	}

	// 测试超时情况可以通过设置一个很短的超时时间来模拟，但这可能会使测试不稳定
}

func TestConvertToInlineImageData(t *testing.T) {
	imageData := []byte("fake image data")
	contentType := "image/jpeg"

	result := ConvertToInlineImageData(imageData, contentType)
	expected := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imageData)

	if result != expected {
		t.Errorf("ConvertToInlineImageData failed. Expected %s, got %s", expected, result)
	}
}

func TestInferImageType(t *testing.T) {
	testCases := []struct {
		name     string
		data     []byte
		expected string
	}{
		{"PNG", []byte("\x89PNG\r\n\x1a\n"), "image/png"},
		{"WEBP", []byte("RIFF\x00\x00\x00\x00WEBP"), "image/webp"},
		{"Unknown", []byte("unknown data"), "application/octet-stream"},
		{"Empty", []byte{}, "application/octet-stream"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := InferImageType(tc.data)
			if result != tc.expected {
				t.Errorf("InferImageType failed for %s. Expected %s, got %s", tc.name, tc.expected, result)
			}
		})
	}
}

func TestURLToInlineImageData(t *testing.T) {
	url := "https://profile.line-scdn.net/0hAzoguxL4HlVSHg3Kc2FhAm5bEDglMBgdKitSOiUaEjJ3JlwFPitZMHMWQWwqKw5XZ34GM3MfRGB2"
	result, err := URLToInlineImageData(url)
	if err != nil {
		t.Errorf("URLToInlineImageData failed: %v", err)
	}
	t.Logf("result: %s", result)
}
