package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ddgf-new/internal/model"
)

func (r *mutationResolver) CreateAdvice(ctx context.Context, content string) (*model.Advice, error) {
	advice, _ := r.DB.CreateAdvice(content)
	return advice.ToModel(), nil
}

func (r *mutationResolver) DeleteAdvice(ctx context.Context, id string) (bool, error) {
	err := r.DB.DeleteAdvice(id)
	return err == nil, err
}

func (r *mutationResolver) Tag(ctx context.Context, tid string, aid string) (*model.Advice, error) {
	advice, err := r.DB.TagAdvice(tid, aid)
	if err != nil {
		return nil, err
	}
	return advice.ToModel(), nil
}

func (r *mutationResolver) Untag(ctx context.Context, tid string, aid string) (*model.Advice, error) {
	advice, err := r.DB.UntagAdvice(tid, aid)
	if err != nil {
		return nil, err
	}
	return advice.ToModel(), err
}

func (r *queryResolver) Advices(ctx context.Context) ([]*model.Advice, error) {
	mAdvices := make([]*model.Advice, 0)
	advices, err := r.DB.GetAdvices()
	if err != nil {
		return mAdvices, err
	}
	for _, advice := range advices {
		mAdvices = append(mAdvices, advice.ToModel())
	}
	return mAdvices, nil
}

func (r *queryResolver) Advice(ctx context.Context, id string) (*model.Advice, error) {
	advice, err := r.DB.GetAdvice(id)
	if err != nil {
		return nil, err
	}
	return advice.ToModel(), nil
}
