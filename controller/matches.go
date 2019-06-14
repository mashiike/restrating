package controller

import (
	"context"
	"log"

	matches "github.com/mashiike/restrating/gen/matches"
	"github.com/mashiike/restrating/usecase"
)

// Matches service example implementation.
// The example methods log the requests and return zero values.
type matchessrvc struct {
	logger *log.Logger
	uc     *usecase.CoreUsecase
}

// NewMatches returns the Matches service implementation.
func NewMatches(logger *log.Logger, uc *usecase.CoreUsecase) matches.Service {
	return &matchessrvc{logger, uc}
}

// Apply match and return RatingResources
func (s *matchessrvc) ApplyMatch(ctx context.Context, p *matches.ApplyMatchPayload) (res *matches.RestratingMatch, err error) {
	s.logger.Print("matches.Apply match")
	var input *ApplyMatchInput
	input, err = newApplyMatchInput(p)
	if err != nil {
		return
	}
	output := newApplyMatchOutput()
	err = s.uc.ApplyMatch(ctx, input, output)
	if err == nil {
		res = output.response
	}
	return
}
