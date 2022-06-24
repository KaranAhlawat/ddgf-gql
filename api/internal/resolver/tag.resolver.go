package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ddgf-new/internal/model"
)

func (r *mutationResolver) CreateTag(ctx context.Context, tag string) (*model.Tag, error) {
	retTag, _ := r.DB.CreateTag(tag)
	return retTag.ToModel(), nil
}

func (r *mutationResolver) DeleteTag(ctx context.Context, id string) (bool, error) {
	err := r.DB.DeleteTag(id)
	return err == nil, err
}

func (r *queryResolver) Tags(ctx context.Context) ([]*model.Tag, error) {
	mTags := make([]*model.Tag, 0)
	tags, err := r.DB.GetTags()
	if err != nil {
		return mTags, err
	}
	for _, tag := range tags {
		mTags = append(mTags, tag.ToModel())
	}
	return mTags, nil
}

func (r *queryResolver) Tag(ctx context.Context, id string) (*model.Tag, error) {
	tag, err := r.DB.GetTag(id)
	if err != nil {
		return nil, err
	}
	return tag.ToModel(), nil
}

func (r *queryResolver) AdvicesForTag(ctx context.Context, id string) ([]*model.Advice, error) {
	mAdvices := make([]*model.Advice, 0)
	advices, err := r.DB.AdvicesForTag(id)
	if err != nil {
		return mAdvices, err
	}
	for _, advice := range advices {
		mAdvices = append(mAdvices, advice.ToModel())
	}
	return mAdvices, nil
}
