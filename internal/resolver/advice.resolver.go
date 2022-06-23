package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ddgf-new/internal/model"
	"fmt"
)

func (r *mutationResolver) CreateAdvice(ctx context.Context, content string) (*model.Advice, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAdvice(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Advices(ctx context.Context) ([]*model.Advice, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Advice(ctx context.Context, id string) (*model.Advice, error) {
	panic(fmt.Errorf("not implemented"))
}
