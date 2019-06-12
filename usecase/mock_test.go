package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/mashiike/restrating/entity"
	"github.com/mashiike/restrating/usecase"
)

type fakeRatingResourceRepository struct {
	usecase.RatingResourceRepository
	t                 *testing.T
	FakeFindAllByRRNs func(*testing.T, []entity.RRN) (entity.RatingResourceCollection, error)
	FakeSave          func(*testing.T, entity.RatingResource) error
	FakeSaveAll       func(*testing.T, entity.RatingResourceCollection) error
}

func (r *fakeRatingResourceRepository) FindAllByRRNs(_ context.Context, rrns []entity.RRN) (entity.RatingResourceCollection, error) {
	if r.FakeFindAllByRRNs == nil {
		r.t.Fatal("unexpected RatingResourceRepository.FindAllByRRNs call")
	}
	return r.FakeFindAllByRRNs(r.t, rrns)
}

func (r *fakeRatingResourceRepository) Save(_ context.Context, resouce entity.RatingResource) error {
	if r.FakeSave == nil {
		r.t.Fatal("unexpected RatingResourceRepository.Save call")
	}
	return r.FakeSave(r.t, resouce)
}

func (r *fakeRatingResourceRepository) SaveAll(ctx context.Context, resouces entity.RatingResourceCollection) error {
	if r.FakeSaveAll == nil {
		r.t.Fatal("unexpected RatingResourceRepository.SaveAll call")
	}
	return r.FakeSaveAll(r.t, resouces)
}

type fakeApplyMatchInput struct {
	usecase.ApplyMatchInput
	scores    entity.ScoreCollection
	outcomeAt time.Time
}

func (i *fakeApplyMatchInput) Scores() entity.ScoreCollection {
	return i.scores
}

func (i *fakeApplyMatchInput) OutcomeAt() time.Time {
	return i.outcomeAt
}

type fakeApplyMatchOutput struct {
	usecase.ApplyMatchOutput
	t                   *testing.T
	FakeSetParticipants func(*testing.T, entity.RatingResourceCollection)
}

func (o *fakeApplyMatchOutput) SetParticipants(collection entity.RatingResourceCollection) {
	if o.FakeSetParticipants == nil {
		o.t.Fatal("unexpected ApplyMatchOutput.SetParticipants call")
	}
	o.FakeSetParticipants(o.t, collection)
}

type fakeCreatePlayerInput struct {
	usecase.CreatePlayerInput
	name string
	now  time.Time
}

func (i *fakeCreatePlayerInput) Name() string {
	return i.name
}

func (i *fakeCreatePlayerInput) Now() time.Time {
	return i.now
}

type fakeCreatePlayerOutput struct {
	t             *testing.T
	FakeSetPlayer func(*testing.T, *entity.Player)
}

func (o *fakeCreatePlayerOutput) SetPlayer(p *entity.Player) {
	if o.FakeSetPlayer == nil {
		o.t.Fatal("unexpected CreatePlayerOutput.SetPlayer call")
	}
	o.FakeSetPlayer(o.t, p)
}
