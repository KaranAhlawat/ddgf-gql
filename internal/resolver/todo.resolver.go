package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ddgf-new/internal/model"
	"fmt"
	"math/rand"
)

func (r *mutationResolver) CreateTodo(_ context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   input.UserID,
			Name: "User" + input.UserID,
		},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(_ context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}
