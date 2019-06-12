package usecase

import (
	"context"
	"time"

	"github.com/mashiike/restrating/entity"
	"github.com/pkg/errors"
)

//CoreUsecase はアプリケーションの振る舞いを記述する。
type CoreUsecase struct {
	config    *entity.Config
	repo      RatingResourceRepository
	lifeCycle LifeCycle
}

//New はCoreUsecaseのコンストラクタ
func New(config *entity.Config, repo RatingResourceRepository, lifeCycle LifeCycle) *CoreUsecase {
	return &CoreUsecase{
		config:    config,
		repo:      repo,
		lifeCycle: lifeCycle,
	}
}

//RatingResourceRepository はRatingResouce(Player/Team/etc...) にアクセスするための永続化インタフェース
type RatingResourceRepository interface {
	FindAllByRRNs(context.Context, []entity.RRN) (entity.RatingResourceCollection, error)
	Save(context.Context, entity.RatingResource) error
	SaveAll(context.Context, entity.RatingResourceCollection) error
}

//LifeCycle はアプリケーションのライフサイクルを管理します
type LifeCycle interface {
	Begin(context.Context) context.Context
	End(context.Context, error)
}

//CreatePlayerInput はCreatePlayer usecase のInputポート
type CreatePlayerInput interface {
	Name() string
	Now() time.Time
}

//CreatePlayerInput はCreatePlayer usecase のOutputポート
type CreatePlayerOutput interface {
	SetPlayer(*entity.Player)
}

//CreatePlayer はPlayerというRatingResouceを追加します。
func (u *CoreUsecase) CreatePlayer(ctx context.Context, input CreatePlayerInput, output CreatePlayerOutput) (err error) {
	ctx = u.lifeCycle.Begin(ctx)
	defer func() { u.lifeCycle.End(ctx, err) }()

	player := entity.NewPlayer(input.Name(), input.Now(), u.config)
	if e := u.repo.Save(ctx, player); e != nil {
		err = errors.Wrap(e, "save player")
		return
	}
	output.SetPlayer(player)
	return
}

//ApplyMatchInput　はApplyMatch usecase のInputポート
type ApplyMatchInput interface {
	Scores() entity.ScoreCollection
	OutcomeAt() time.Time
}

//ApplyMatchInput　はApplyMatch usecase のOutputポート
type ApplyMatchOutput interface {
	SetParticipants(entity.RatingResourceCollection)
}

//ApplyMatch はRatingResouce同士の対戦をそれぞれのRatingに反映します。
func (u *CoreUsecase) ApplyMatch(ctx context.Context, input ApplyMatchInput, output ApplyMatchOutput) (err error) {
	ctx = u.lifeCycle.Begin(ctx)
	defer func() { u.lifeCycle.End(ctx, err) }()

	scores := input.Scores()
	participants, e := u.repo.FindAllByRRNs(ctx, scores.Keys())
	if e != nil {
		err = errors.Wrap(e, "find")
		return
	}
	if e := participants.Prepares(input.OutcomeAt(), u.config); e != nil {
		err = errors.Wrap(e, "prepare")
		return
	}
	if e := participants.Updates(scores, input.OutcomeAt(), u.config); e != nil {
		err = errors.Wrap(e, "update")
		return
	}

	if e := u.repo.SaveAll(ctx, participants); e != nil {
		err = errors.Wrap(e, "save")
		return
	}
	output.SetParticipants(participants)
	return
}
