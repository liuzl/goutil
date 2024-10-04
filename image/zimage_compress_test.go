package zimage

import (
	"testing"
)

func TestAutoCompressEdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		{"Empty Data", []byte{}, true},
		{"Very Small Data", []byte{0, 1, 2, 3}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compressed, err := AutoCompress(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("AutoCompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && len(compressed) == 0 {
				t.Errorf("AutoCompress() returned empty data for non-error case")
			}
		})
	}
}
