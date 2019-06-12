package entity

import (
	"sort"
	"time"

	"github.com/mashiike/rating"
	"github.com/pkg/errors"
)

//RatingResource はPlayer/Teamの同一視
type RatingResource interface {
	RRN() RRN
	Rating() rating.Rating
	Prepare(time.Time, *Config) error
	Update(*MatchResult) error
}

//RatingResourceCollection はRatingResouceのコレクション
type RatingResourceCollection map[RRN]RatingResource

//Add はRatingResourceCollectionに一気に追加する関数
func (c RatingResourceCollection) Add(items ...RatingResource) RatingResourceCollection {
	for _, item := range items {
		c[item.RRN()] = item
	}
	return c
}

//Ratings はその時点でのRatingの値を返す
func (c RatingResourceCollection) Ratings() map[RRN]rating.Rating {
	ratings := make(map[RRN]rating.Rating, len(c))
	for rrn, resource := range c {
		ratings[rrn] = resource.Rating()
	}
	return ratings
}

//Prepares はまとめてコレクションをPrepareする。
func (c RatingResourceCollection) Prepares(outcomeAt time.Time, config *Config) error {

	for rrn, resource := range c {
		if err := resource.Prepare(outcomeAt, config); err != nil {
			return errors.Wrapf(err, "rrn=%s", rrn)
		}
	}
	return nil
}

//ScoreCollection はまとめてアップデートするためのScoreの集合
type ScoreCollection map[RRN]float64

//Keys はScoreの集合のKeyになるRRNのスライス
func (c ScoreCollection) Keys() []RRN {
	rrns := make([]RRN, 0, len(c))
	for rrn, _ := range c {
		rrns = append(rrns, rrn)
	}
	return rrns
}

//Values はRatingResouceのスライスを返します。
func (c RatingResourceCollection) Values() []RatingResource {
	resources := make([]RatingResource, 0, len(c))
	for _, resource := range c {
		resources = append(resources, resource)
	}
	return resources
}

//Updates はまとめてコレクションをUpdateする。
func (c RatingResourceCollection) Updates(scores ScoreCollection, outcomeAt time.Time, config *Config) error {

	snapshots := c.Ratings()
	for rrn1, resource := range c {
		score1 := scores[rrn1]
		for rrn2, score2 := range scores {
			if rrn1 == rrn2 {
				continue
			}
			s := rating.ScoreWin
			if score1 < score2 {
				s = rating.ScoreLose
			}
			if score1 == score2 {
				s = rating.ScoreDraw
			}
			err := resource.Update(&MatchResult{
				Opponent:  snapshots[rrn2],
				Score:     s,
				OutcomeAt: outcomeAt,
			})
			if err != nil {
				return errors.Wrapf(err, "update %s", rrn1)
			}
		}
	}
	return nil
}

//Ranking は強さ順に並んだSliceを返す
func (c RatingResourceCollection) Ranking() []RatingResource {
	resources := c.Values()
	ratings := c.Ratings()
	sort.Slice(resources, func(i, j int) bool {
		ratingi := ratings[resources[i].RRN()]
		ratingj := ratings[resources[j].RRN()]
		if ratingi.IsWeeker(ratingj) {
			return false
		}
		return ratingi.Strength() > ratingj.Strength()
	})
	return resources
}
