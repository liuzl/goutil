package rest

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/justinas/alice"
	"github.com/liuzl/filestore"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

var (
	dir        = flag.String("zlog_dir", filepath.Join(filepath.Dir(os.Args[0]), "zerolog"), "zerolog dir")
	alsoStdout = flag.Bool("zlog_stdout", false, "also output logs to stdout")
)

var c alice.Chain
var once sync.Once
var log zerolog.Logger

func Log() *zerolog.Logger {
	initZlog()
	return &log
}

func initZlog() {
	once.Do(func() {
		zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
			return filepath.Base(file) + ":" + strconv.Itoa(line)
		}
		hostname, _ := os.Hostname()
		var out io.Writer
		f, err := filestore.NewFileStore(*dir)
		if err != nil {
			out = zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false}
			fmt.Fprintf(os.Stderr, "err: %+v, will zerolog to stdout\n", err)
		} else {
			if *alsoStdout {
				consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false}
				out = io.MultiWriter(f, consoleWriter)
			} else {
				out = f
			}
		}
		log = zerolog.New(out).With().
			Timestamp().
			Caller().
			Str("service", filepath.Base(os.Args[0])).
			Str("host", hostname).
			Logger()

		c = alice.New()

		// Install the logger handler with default output on the console
		c = c.Append(hlog.NewHandler(log))

		c = c.Append(hlog.AccessHandler(func(r *http.Request,
			status, size int, duration time.Duration) {
			hlog.FromRequest(r).Debug().
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
		c = c.Append(HeaderHandler("X-Forwarded-For"))
		c = c.Append(HeaderHandler("User-Agent"))
		c = c.Append(HeaderHandler("Referer"))
		c = c.Append(RequestIDHandler("req_id", "Request-Id"))
		c = c.Append(DumpResponseHandler("response"))
		c = c.Append(DumpRequestHandler("request"))
	})
}

func WithLog(f func(http.ResponseWriter, *http.Request)) http.Handler {
	initZlog()
	return c.Then(http.HandlerFunc(f))
}
