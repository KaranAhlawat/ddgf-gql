package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ddgf-new/internal/model"
	"fmt"
)

func (r *mutationResolver) CreatePage(ctx context.Context, content string) (*model.Page, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePage(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Pages(ctx context.Context) ([]*model.Page, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Page(ctx context.Context, id string) (*model.Page, error) {
	panic(fmt.Errorf("not implemented"))
}
