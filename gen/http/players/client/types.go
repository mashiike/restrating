// Code generated by goa v3.0.2, DO NOT EDIT.
//
// Players HTTP client types
//
// Command:
// $ goa gen github.com/mashiike/restrating/design

package client

import (
	players "github.com/mashiike/restrating/gen/players"
	playersviews "github.com/mashiike/restrating/gen/players/views"
)

// CreatePlayerRequestBody is the type of the "Players" service "create player"
// endpoint HTTP request body.
type CreatePlayerRequestBody struct {
	Name string `form:"name" json:"name" xml:"name"`
}

// CreatePlayerResponseBody is the type of the "Players" service "create
// player" endpoint HTTP response body.
type CreatePlayerResponseBody struct {
	// Rating Resource Name
	Rrn *string `form:"rrn,omitempty" json:"rrn,omitempty" xml:"rrn,omitempty"`
}

// NewCreatePlayerRequestBody builds the HTTP request body from the payload of
// the "create player" endpoint of the "Players" service.
func NewCreatePlayerRequestBody(p *players.CreatePlayerPayload) *CreatePlayerRequestBody {
	body := &CreatePlayerRequestBody{
		Name: p.Name,
	}
	return body
}

// NewCreatePlayerRestratingRrnCreated builds a "Players" service "create
// player" endpoint result from a HTTP "Created" response.
func NewCreatePlayerRestratingRrnCreated(body *CreatePlayerResponseBody) *playersviews.RestratingRrnView {
	v := &playersviews.RestratingRrnView{
		Rrn: body.Rrn,
	}
	return v
}