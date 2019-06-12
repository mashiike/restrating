package entity_test

import (
	"errors"
	"testing"

	"github.com/mashiike/restrating/entity"
)

func TestParseRRN(t *testing.T) {
	cases := []struct {
		input string
		rrn   entity.RRN
		err   error
	}{
		{
			input: "invalid",
			err:   errors.New("rrn: invalid prefix"),
		},
		{
			input: "rrn:org",
			err:   errors.New("rrn: not enough sections"),
		},
		{
			input: "rrn:team:goat",
			rrn: entity.RRN{
				Type: "team",
				Name: "goat",
			},
		},
	}
	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			got, err := entity.ParseRRN(c.input)
			if c.rrn != got {
				t.Errorf("rrn.Parse(%v) unexpected %v, expected %v", c.input, got, c.rrn)
			}
			if err == nil && c.err != nil {
				t.Errorf("expected err %v, but got nil", c.err)
			} else if err != nil && c.err == nil {
				t.Errorf("expected err nil, but got %v", err)
			} else if err != nil && c.err != nil && err.Error() != c.err.Error() {
				t.Errorf("expected err %v, but got %v", c.err, err)
			}
		})
	}
}
