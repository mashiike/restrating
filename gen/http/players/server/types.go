// Code generated by goa v3.0.2, DO NOT EDIT.
//
// Players HTTP server types
//
// Command:
// $ goa gen github.com/mashiike/restrating/design

package server

import (
	players "github.com/mashiike/restrating/gen/players"
	playersviews "github.com/mashiike/restrating/gen/players/views"
	goa "goa.design/goa/v3/pkg"
)

// CreatePlayerRequestBody is the type of the "Players" service "create player"
// endpoint HTTP request body.
type CreatePlayerRequestBody struct {
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// CreatePlayerResponseBody is the type of the "Players" service "create
// player" endpoint HTTP response body.
type CreatePlayerResponseBody struct {
	// Rating Resource Name
	Rrn string `form:"rrn" json:"rrn" xml:"rrn"`
}

// NewCreatePlayerResponseBody builds the HTTP response body from the result of
// the "create player" endpoint of the "Players" service.
func NewCreatePlayerResponseBody(res *playersviews.RestratingRrnView) *CreatePlayerResponseBody {
	body := &CreatePlayerResponseBody{
		Rrn: *res.Rrn,
	}
	return body
}

// NewCreatePlayerPayload builds a Players service create player endpoint
// payload.
func NewCreatePlayerPayload(body *CreatePlayerRequestBody) *players.CreatePlayerPayload {
	v := &players.CreatePlayerPayload{
		Name: *body.Name,
	}
	return v
}

// ValidateCreatePlayerRequestBody runs the validations defined on Create
// PlayerRequestBody
func ValidateCreatePlayerRequestBody(body *CreatePlayerRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}
