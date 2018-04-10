package rest

import (
	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var c alice.Chain

func init() {
	hostname, _ := os.Hostname()
	log := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("service", filepath.Base(os.Args[0])).
		Str("host", hostname).
		Logger()

	c = alice.New()

	// Install the logger handler with default output on the console
	c = c.Append(hlog.NewHandler(log))

	c = c.Append(hlog.AccessHandler(func(r *http.Request,
		status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))

	// Install some provided extra handler to set some request's context fields.
	// Thanks to those handler, all our logs will come with some pre-populated fields.
	c = c.Append(hlog.RemoteAddrHandler("ip"))
	c = c.Append(hlog.UserAgentHandler("user_agent"))
	c = c.Append(hlog.RefererHandler("referer"))
	c = c.Append(hlog.RequestIDHandler("req_id", "Request-Id"))
}

func WithLog(f func(http.ResponseWriter, *http.Request)) http.Handler {
	return c.Then(http.HandlerFunc(f))
}
