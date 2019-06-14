package inmem

import (
	"context"
	"fmt"

	"github.com/mashiike/restrating/entity"
)

type RatingResourceRepository struct {
	strage entity.RatingResourceCollection
}

func NewRatingResourceRepository() *RatingResourceRepository {
	return &RatingResourceRepository{
		strage: make(entity.RatingResourceCollection, 10),
	}
}

func (r *RatingResourceRepository) FindAllByRRNs(_ context.Context, rrns []entity.RRN) (entity.RatingResourceCollection, error) {
	ret := make(entity.RatingResourceCollection, len(rrns))
	for _, rrn := range rrns {
		resouce, ok := r.strage[rrn]
		if !ok {
			return nil, fmt.Errorf("not found rrn = %s", rrn)
		}
		ret[rrn] = resouce
	}
	return ret, nil
}

func (r *RatingResourceRepository) Save(ctx context.Context, resouce entity.RatingResource) error {
	return r.SaveAll(ctx, entity.RatingResourceCollection{resouce.RRN(): resouce})
}

func (r *RatingResourceRepository) SaveAll(_ context.Context, resources entity.RatingResourceCollection) error {
	for rrn, resouce := range resources {
		r.strage[rrn] = resouce
	}
	return nil
}
