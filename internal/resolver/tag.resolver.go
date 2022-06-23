package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ddgf-new/internal/model"
	"fmt"
)

func (r *mutationResolver) CreateTag(ctx context.Context, tag string) (*model.Tag, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTag(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tags(ctx context.Context) ([]*model.Tag, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tag(ctx context.Context, id string) (*model.Tag, error) {
	panic(fmt.Errorf("not implemented"))
}
