// Code generated by goa v3.0.2, DO NOT EDIT.
//
// Players HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/mashiike/restrating/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	players "github.com/mashiike/restrating/gen/players"
	playersviews "github.com/mashiike/restrating/gen/players/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildCreatePlayerRequest instantiates a HTTP request object with method and
// path set to call the "Players" service "create player" endpoint
func (c *Client) BuildCreatePlayerRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreatePlayerPlayersPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Players", "create player", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreatePlayerRequest returns an encoder for requests sent to the
// Players create player server.
func EncodeCreatePlayerRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*players.CreatePlayerPayload)
		if !ok {
			return goahttp.ErrInvalidType("Players", "create player", "*players.CreatePlayerPayload", v)
		}
		body := NewCreatePlayerRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("Players", "create player", err)
		}
		return nil
	}
}

// DecodeCreatePlayerResponse returns a decoder for responses returned by the
// Players create player endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeCreatePlayerResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreatePlayerResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Players", "create player", err)
			}
			p := NewCreatePlayerRestratingRrnCreated(&body)
			view := "default"
			vres := &playersviews.RestratingRrn{p, view}
			if err = playersviews.ValidateRestratingRrn(vres); err != nil {
				return nil, goahttp.ErrValidationError("Players", "create player", err)
			}
			res := players.NewRestratingRrn(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Players", "create player", resp.StatusCode, string(body))
		}
	}
}