package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ddgf-new/internal/graph"
)

func (r *mutationResolver) Ping(_ context.Context, text string) (string, error) {
	return "Pong " + text, nil
}

func (r *queryResolver) Ping(_ context.Context) (string, error) {
	return "Pong", nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
