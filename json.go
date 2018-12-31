package goutil

import (
	"bytes"
	"encoding/json"
)

// JSONMarshal marshals t in json without HTML escaping
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

// JSONMarshalIndent marshals t in json with prefix and indention
func JSONMarshalIndent(t interface{}, prefix, indent string) ([]byte, error) {
	b, err := JSONMarshal(t)
	if err != nil {
		return b, err
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, prefix, indent)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
