package controller

import (
	"time"

	"github.com/mashiike/restrating/entity"
	matches "github.com/mashiike/restrating/gen/matches"
	players "github.com/mashiike/restrating/gen/players"
)

var loc *time.Location

func init() {
	loc, _ = time.LoadLocation("Asia/Tokyo")
}

type CreatePlayerInput struct {
	payload *players.CreatePlayerPayload
}

func newCreatePlayerInput(payload *players.CreatePlayerPayload) *CreatePlayerInput {
	return &CreatePlayerInput{
		payload: payload,
	}
}

func (i *CreatePlayerInput) Name() string {
	return i.payload.Name
}
func (_ *CreatePlayerInput) Now() time.Time {
	return time.Now().In(loc)
}

type CreatePlayerOutput struct {
	response *players.RestratingRrn
}

func newCreatePlayerOutput() *CreatePlayerOutput {
	return &CreatePlayerOutput{
		response: &players.RestratingRrn{},
	}
}

func (o *CreatePlayerOutput) SetPlayer(player *entity.Player) {
	o.response.Rrn = player.RRN().String()
}

type ApplyMatchInput struct {
	scores    entity.ScoreCollection
	outcomeAt time.Time
}

func newApplyMatchInput(payload *matches.ApplyMatchPayload) (*ApplyMatchInput, error) {
	scores := make(entity.ScoreCollection, len(payload.Scores))
	for rrnStr, score := range payload.Scores {
		rrn, err := entity.ParseRRN(rrnStr)
		if err != nil {
			return nil, err
		}
		scores[rrn] = score
	}
	return &ApplyMatchInput{
		scores:    scores,
		outcomeAt: time.Now().In(loc),
	}, nil
}

func (i *ApplyMatchInput) Scores() entity.ScoreCollection {
	return i.scores
}

func (i *ApplyMatchInput) OutcomeAt() time.Time {
	return i.outcomeAt
}

type ApplyMatchOutput struct {
	response *matches.RestratingMatch
}

func newApplyMatchOutput() *ApplyMatchOutput {
	return &ApplyMatchOutput{
		response: &matches.RestratingMatch{},
	}
}

func (o *ApplyMatchOutput) SetParticipants(participants entity.RatingResourceCollection) {
	tmp := make([]*matches.RatingResource, 0, len(participants))
	for rrn, p := range participants {
		r := p.Rating()
		lower, upper := r.Interval()
		tmp = append(tmp, &matches.RatingResource{
			Rrn: rrn.String(),
			Rating: &matches.Rating{
				Strength: r.Strength(),
				Lower:    lower,
				Upper:    upper,
			},
		})
	}
	o.response.Participants = tmp
}
