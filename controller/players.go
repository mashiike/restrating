package controller

import (
	"context"
	"log"

	players "github.com/mashiike/restrating/gen/players"
	"github.com/mashiike/restrating/usecase"
)

// Players service example implementation.
// The example methods log the requests and return zero values.
type playerssrvc struct {
	logger *log.Logger
	uc     *usecase.CoreUsecase
}

// NewPlayers returns the Players service implementation.
func NewPlayers(logger *log.Logger, uc *usecase.CoreUsecase) players.Service {
	return &playerssrvc{logger, uc}
}

// Add new player and return its RRN(Rating Resource Name).
func (s *playerssrvc) CreatePlayer(ctx context.Context, p *players.CreatePlayerPayload) (res *players.RestratingRrn, err error) {
	s.logger.Print("players.create player")
	output := newCreatePlayerOutput()
	err = s.uc.CreatePlayer(ctx, newCreatePlayerInput(p), output)
	if err == nil {
		res = output.response
	}
	return
}
