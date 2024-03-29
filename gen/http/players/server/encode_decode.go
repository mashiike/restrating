// Code generated by goa v3.0.2, DO NOT EDIT.
//
// Players HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/mashiike/restrating/design

package server

import (
	"context"
	"io"
	"net/http"

	playersviews "github.com/mashiike/restrating/gen/players/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeCreatePlayerResponse returns an encoder for responses returned by the
// Players create player endpoint.
func EncodeCreatePlayerResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*playersviews.RestratingRrn)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
		enc := encoder(ctx, w)
		body := NewCreatePlayerResponseBody(res.Projected)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreatePlayerRequest returns a decoder for requests sent to the Players
// create player endpoint.
func DecodeCreatePlayerRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreatePlayerRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreatePlayerRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreatePlayerPayload(&body)

		return payload, nil
	}
}
