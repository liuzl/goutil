package rest

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/carlmjohnson/requests"
)

func Record(rt http.RoundTripper, rawReq *string, rawRes *string, duration *time.Duration) requests.Transport {
	if rt == nil {
		rt = http.DefaultTransport
	}
	return requests.RoundTripFunc(func(req *http.Request) (res *http.Response, err error) {
		start := time.Now()
		defer func() {
			if err != nil {
				err = fmt.Errorf("problem while recording transport: %w", err)
			}
		}()

		b, err := httputil.DumpRequest(req, true)
		if err != nil {
			return nil, err
		}
		*rawReq = string(b)
		if res, err = rt.RoundTrip(req); err != nil {
			return
		}
		*duration = time.Since(start)
		b, err = httputil.DumpResponse(res, true)
		if err != nil {
			return nil, err
		}
		*rawRes = string(b)
		return
	})
}
