package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5(t *testing.T) {
	cases := map[string]string{
		"":       "d41d8cd98f00b204e9800998ecf8427e",
		"123456": "e10adc3949ba59abbe56e057f20f883e",
	}
	for k, v := range cases {
		v1 := MD5(k)
		assert.Equal(t, v, v1, "should be equal")
	}
}

func TestBase62MD5(t *testing.T) {
	cases := map[string]string{
		"":         "JHG80zb3GpRNZk91ZULhYZ",
		"123456":   "KQ2pcfuCXtra1j36xvooAT",
		"password": "8CvFnewHEDjeHJqFEPvCTH",
	}
	for k, v := range cases {
		v1 := Base62MD5(k)
		assert.Equal(t, v, v1, "should be equal")
	}
}
