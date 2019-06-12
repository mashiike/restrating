package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/mashiike/restrating/entity"
	"github.com/mashiike/restrating/usecase"
)

func testBaseTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	baseTime := time.Date(2019, 05, 01, 0, 0, 0, 0, loc)
	return baseTime
}

func TestCoreApplyMatch(t *testing.T) {
	baseTime := testBaseTime()
	config := entity.NewConfig()

	sheepRRN := entity.MustParseRNN("rrn:player:sheep")
	goatRRN := entity.MustParseRNN("rrn:player:goat")

	testCases := []struct {
		name        string
		input       *fakeApplyMatchInput
		output      *fakeApplyMatchOutput
		repo        *fakeRatingResourceRepository
		expectedErr string
	}{
		{
			name: "success case",
			input: &fakeApplyMatchInput{
				scores: entity.ScoreCollection{
					sheepRRN: 1.0,
					goatRRN:  0.0,
				},
				outcomeAt: baseTime.AddDate(0, 0, 8),
			},
			output: &fakeApplyMatchOutput{
				FakeSetParticipants: func(t *testing.T, c entity.RatingResourceCollection) {
					if len(c) != 2 {
						t.Errorf("unexpected collection size = %d", len(c))
						t.Logf("collection is %v", c)
					}

					sheep, ok := c[sheepRRN]
					if !ok {
						t.Error("unexpected output sheep rating not set")
					}
					goat, ok := c[goatRRN]
					if !ok {
						t.Error("unexpected output goat rating not set")
					}
					if goat.Rating().IsWeeker(sheep.Rating()) {
						t.Error("unexpected rating relation")
						t.Logf("goat %v", goat)
						t.Logf("sheep %v", sheep)
					}
				},
			},
			repo: &fakeRatingResourceRepository{
				FakeFindAllByRRNs: func(t *testing.T, rrns []entity.RRN) (entity.RatingResourceCollection, error) {
					if len(rrns) != 2 {
						t.Error("repository input not 2 rrns")
					}
					players := make(entity.RatingResourceCollection, 2)
					for _, rrn := range rrns {
						players[rrn] = entity.NewPlayer(rrn.Name, baseTime, config)
					}
					return players, nil
				},
				FakeSaveAll: func(t *testing.T, c entity.RatingResourceCollection) error {
					if len(c) != 2 {
						t.Error("save target not 2 resouces")
					}
					return nil
				},
			},
		},
		{
			name: "resouce not found",
			input: &fakeApplyMatchInput{
				scores: entity.ScoreCollection{
					sheepRRN: 1.0,
					goatRRN:  0.0,
				},
				outcomeAt: baseTime.AddDate(0, 0, 8),
			},
			output: &fakeApplyMatchOutput{},
			repo: &fakeRatingResourceRepository{
				FakeFindAllByRRNs: func(t *testing.T, rrns []entity.RRN) (entity.RatingResourceCollection, error) {
					return nil, errors.New("resouce not found")
				},
			},
			expectedErr: "find: resouce not found",
		},
		{
			name: "save failed case",
			input: &fakeApplyMatchInput{
				scores: entity.ScoreCollection{
					sheepRRN: 1.0,
					goatRRN:  0.0,
				},
				outcomeAt: baseTime.AddDate(0, 0, 8),
			},
			output: &fakeApplyMatchOutput{},
			repo: &fakeRatingResourceRepository{
				FakeFindAllByRRNs: func(t *testing.T, rrns []entity.RRN) (entity.RatingResourceCollection, error) {
					players := make(entity.RatingResourceCollection, 2)
					for _, rrn := range rrns {
						players[rrn] = entity.NewPlayer(rrn.Name, baseTime, config)
					}
					return players, nil
				},
				FakeSaveAll: func(t *testing.T, c entity.RatingResourceCollection) error {
					return errors.New("unknown")
				},
			},
			expectedErr: "save: unknown",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.repo.t = t
			tc.output.t = t
			lc := newLifeCycle(t)
			u := usecase.New(
				config,
				tc.repo,
				lc,
			)
			if err := u.ApplyMatch(context.Background(), tc.input, tc.output); !checkErr(err, tc.expectedErr) {
				t.Errorf("unexpected err %v, expected %v", err, tc.expectedErr)
			}
			if !checkErr(lc.lastErr, tc.expectedErr) {
				t.Errorf("LifeCycle trap error unexpected %v", lc.lastErr)
			}

		})
	}
}

func TestCoreCreatePlayer(t *testing.T) {
	baseTime := testBaseTime()
	config := entity.NewConfig()

	testCases := []struct {
		name        string
		input       *fakeCreatePlayerInput
		output      *fakeCreatePlayerOutput
		repo        *fakeRatingResourceRepository
		expectedErr string
	}{
		{
			name: "success case",
			input: &fakeCreatePlayerInput{
				name: "sheep",
				now:  baseTime,
			},
			output: &fakeCreatePlayerOutput{
				FakeSetPlayer: func(t *testing.T, p *entity.Player) {
					if p.String() != "rrn:player:sheep 1500.0p-700.0(0/0/0)" {
						t.Errorf("unexpected player status: %v", p)
					}
				},
			},
			repo: &fakeRatingResourceRepository{
				FakeSave: func(t *testing.T, r entity.RatingResource) error {
					p, ok := r.(*entity.Player)
					if !ok {
						t.Errorf("unexpected save target type: %#v", r)
					}
					if p.RRN().String() != "rrn:player:sheep" {
						t.Errorf("unexpected rrn: %v", p.RRN())
					}
					return nil
				},
			},
		},
		{
			name: "duplicate case",
			input: &fakeCreatePlayerInput{
				name: "sheep",
				now:  baseTime,
			},
			output: &fakeCreatePlayerOutput{},
			repo: &fakeRatingResourceRepository{
				FakeSave: func(t *testing.T, r entity.RatingResource) error {
					return errors.New("duplicate entry")
				},
			},
			expectedErr: "save player: duplicate entry",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.repo.t = t
			tc.output.t = t
			lc := newLifeCycle(t)
			u := usecase.New(
				config,
				tc.repo,
				lc,
			)
			if err := u.CreatePlayer(context.Background(), tc.input, tc.output); !checkErr(err, tc.expectedErr) {
				t.Errorf("unexpected err %v, expected %v", err, tc.expectedErr)
			}
			if !checkErr(lc.lastErr, tc.expectedErr) {
				t.Errorf("LifeCycle trap error unexpected %v", lc.lastErr)
			}

		})
	}
}

func checkErr(err error, msg string) bool {
	switch {
	case err == nil && msg == "":
		return true
	case err != nil && err.Error() == msg:
		return true
	}
	return false
}
