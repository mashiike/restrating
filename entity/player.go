package entity

import (
	"time"

	"github.com/mashiike/rating"
	"github.com/pkg/errors"
)

//Player is entity of ratable person.
type Player struct {
	rrn       RRN
	estimated *rating.Estimated
	gameCount GameCount
	fixedAt   time.Time
	updatedAt time.Time
	createdAt time.Time
}

//NewPlayer is Player constractor
func NewPlayer(name string, createdAt time.Time, config *Config) *Player {

	return &Player{
		rrn: RRN{
			Type: "player",
			Name: name,
		},
		estimated: rating.NewEstimated(
			config.DefaultRating(),
			config.Tau,
		),
		fixedAt:   createdAt.Truncate(config.RatingPeriod),
		createdAt: createdAt,
	}
}

//Prepare must do before Update
func (p *Player) Prepare(now time.Time, config *Config) error {
	//Reflects the previous non-match period.
	for now.Sub(p.fixedAt) > config.RatingPeriod {
		if err := p.estimated.Fix(); err != nil {
			return err
		}
		p.fixedAt = p.fixedAt.Add(config.RatingPeriod)
		p.estimated = rating.NewEstimated(p.estimated.Fixed, config.Tau)
	}
	return nil
}

//MatchResult represents the Match result from the viewpoint of a player
type MatchResult struct {
	Opponent  rating.Rating `json:"opponent"`
	Score     float64       `json:"score"`
	OutcomeAt time.Time     `json:"outcome_at"`
}

//Update do Player's rating update.
func (p *Player) Update(result *MatchResult) error {
	if p.fixedAt.After(result.OutcomeAt) {
		return errors.New("a match from the pasted")
	}
	if err := p.estimated.ApplyMatch(result.Opponent, result.Score); err != nil {
		return errors.Wrap(err, "player update")
	}
	p.gameCount = p.gameCount.AddScore(result.Score)
	if p.updatedAt.Before(result.OutcomeAt) {
		p.updatedAt = result.OutcomeAt
	}
	return nil
}

//Rating is current this player's estimated rating
func (p *Player) Rating() rating.Rating {
	return p.estimated.Rating()
}

//GameCount is player's game count (win/lose/draw)
func (p *Player) GameCount() GameCount {
	return p.gameCount
}

//RRN はリソースを管理する名前
func (p *Player) RRN() RRN {
	return p.rrn
}

//String is implements fmt.Stringer
func (p *Player) String() string {
	b := make([]byte, 0, 30+len(rating.PlusMinusFormat)+p.gameCount.byteLength())
	b = append(b, p.rrn.String()...)
	b = append(b, ' ')
	b = p.Rating().AppendFormat(b, rating.PlusMinusFormat)
	b = p.gameCount.appendByte(b)
	return string(b)
}
