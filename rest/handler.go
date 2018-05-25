package rest

import (
	"context"
	"net/http"
	"net/http/httputil"

	"github.com/rs/xid"
	"github.com/rs/zerolog"
)

type idKey struct{}

// IDFromRequest returns the unique id associated to the request if any.
func IDFromRequest(r *http.Request) (id xid.ID, ok bool) {
	if r == nil {
		return
	}
	id, ok = r.Context().Value(idKey{}).(xid.ID)
	return
}

// RequestIDHandler returns a handler setting a unique id to the request which can
// be gathered using IDFromRequest(req). This generated id is added as a field to the
// logger using the passed fieldKey as field name. The id is also added as a response
// header if the headerName is not empty.
//
// The generated id is a URL safe base64 encoded mongo object-id-like unique id.
// Mongo unique id generation algorithm has been selected as a trade-off between
// size and ease of use: UUID is less space efficient and snowflake requires machine
// configuration.
func RequestIDHandler(fieldKey, headerName string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			id, ok := IDFromRequest(r)
			if !ok {
				id = xid.New()
				ctx = context.WithValue(ctx, idKey{}, id)
				r = r.WithContext(ctx)
			}
			if fieldKey != "" {
				log := zerolog.Ctx(ctx)
				log.UpdateContext(func(c zerolog.Context) zerolog.Context {
					return c.Str(fieldKey, id.String())
				})
			}
			if headerName != "" {
				r.Header.Set(headerName, id.String())
			}
			next.ServeHTTP(w, r)
		})
	}
}

func DumpRequestHandler(fieldKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := zerolog.Ctx(r.Context())
			log.UpdateContext(func(c zerolog.Context) zerolog.Context {
				res, err := httputil.DumpRequest(r, true)
				var msg string
				if err != nil {
					msg = err.Error()
				} else {
					msg = string(res)
				}
				return c.Str(fieldKey, msg)
			})
			next.ServeHTTP(w, r)
		})
	}
}
