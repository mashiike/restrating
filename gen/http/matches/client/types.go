// Code generated by goa v3.0.2, DO NOT EDIT.
//
// Matches HTTP client types
//
// Command:
// $ goa gen github.com/mashiike/restrating/design

package client

import (
	matches "github.com/mashiike/restrating/gen/matches"
	matchesviews "github.com/mashiike/restrating/gen/matches/views"
	goa "goa.design/goa/v3/pkg"
)

// ApplyMatchRequestBody is the type of the "Matches" service "Apply match"
// endpoint HTTP request body.
type ApplyMatchRequestBody struct {
	Scores map[string]float64 `form:"scores" json:"scores" xml:"scores"`
}

// ApplyMatchResponseBody is the type of the "Matches" service "Apply match"
// endpoint HTTP response body.
type ApplyMatchResponseBody struct {
	Participants []*RatingResourceResponseBody `form:"participants,omitempty" json:"participants,omitempty" xml:"participants,omitempty"`
}

// RatingResourceResponseBody is used to define fields on response body types.
type RatingResourceResponseBody struct {
	// Rating Resource Name
	Rrn *string `form:"rrn,omitempty" json:"rrn,omitempty" xml:"rrn,omitempty"`
	// players strength
	Rating *RatingResponseBody `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// RatingResponseBody is used to define fields on response body types.
type RatingResponseBody struct {
	Strength *float64 `form:"strength,omitempty" json:"strength,omitempty" xml:"strength,omitempty"`
	Lower    *float64 `form:"lower,omitempty" json:"lower,omitempty" xml:"lower,omitempty"`
	Upper    *float64 `form:"upper,omitempty" json:"upper,omitempty" xml:"upper,omitempty"`
}

// NewApplyMatchRequestBody builds the HTTP request body from the payload of
// the "Apply match" endpoint of the "Matches" service.
func NewApplyMatchRequestBody(p *matches.ApplyMatchPayload) *ApplyMatchRequestBody {
	body := &ApplyMatchRequestBody{}
	if p.Scores != nil {
		body.Scores = make(map[string]float64, len(p.Scores))
		for key, val := range p.Scores {
			tk := key
			tv := val
			body.Scores[tk] = tv
		}
	}
	return body
}

// NewApplyMatchRestratingMatchOK builds a "Matches" service "Apply match"
// endpoint result from a HTTP "OK" response.
func NewApplyMatchRestratingMatchOK(body *ApplyMatchResponseBody) *matchesviews.RestratingMatchView {
	v := &matchesviews.RestratingMatchView{}
	v.Participants = make([]*matchesviews.RatingResourceView, len(body.Participants))
	for i, val := range body.Participants {
		v.Participants[i] = unmarshalRatingResourceResponseBodyToMatchesviewsRatingResourceView(val)
	}
	return v
}

// ValidateRatingResourceResponseBody runs the validations defined on
// RatingResourceResponseBody
func ValidateRatingResourceResponseBody(body *RatingResourceResponseBody) (err error) {
	if body.Rrn == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("rrn", "body"))
	}
	if body.Rating == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("rating", "body"))
	}
	if body.Rating != nil {
		if err2 := ValidateRatingResponseBody(body.Rating); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateRatingResponseBody runs the validations defined on ratingResponseBody
func ValidateRatingResponseBody(body *RatingResponseBody) (err error) {
	if body.Strength == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("strength", "body"))
	}
	if body.Lower == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("lower", "body"))
	}
	if body.Upper == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("upper", "body"))
	}
	return
}
