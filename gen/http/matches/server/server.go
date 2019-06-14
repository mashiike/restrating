// Code generated by goa v3.0.2, DO NOT EDIT.
//
// Matches HTTP server
//
// Command:
// $ goa gen github.com/mashiike/restrating/design

package server

import (
	"context"
	"net/http"

	matches "github.com/mashiike/restrating/gen/matches"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the Matches service endpoint HTTP handlers.
type Server struct {
	Mounts     []*MountPoint
	ApplyMatch http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the Matches service endpoints.
func New(
	e *matches.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"ApplyMatch", "POST", "/v1/matches"},
		},
		ApplyMatch: NewApplyMatchHandler(e.ApplyMatch, mux, dec, enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "Matches" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.ApplyMatch = m(s.ApplyMatch)
}

// Mount configures the mux to serve the Matches endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountApplyMatchHandler(mux, h.ApplyMatch)
}

// MountApplyMatchHandler configures the mux to serve the "Matches" service
// "Apply match" endpoint.
func MountApplyMatchHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/matches", f)
}

// NewApplyMatchHandler creates a HTTP handler which loads the HTTP request and
// calls the "Matches" service "Apply match" endpoint.
func NewApplyMatchHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeApplyMatchRequest(mux, dec)
		encodeResponse = EncodeApplyMatchResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Apply match")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Matches")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}