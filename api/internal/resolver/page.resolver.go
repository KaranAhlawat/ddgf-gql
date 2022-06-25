package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ddgf-new/internal/model"
	u "ddgf-new/internal/util"
)

func (r *mutationResolver) CreatePage(ctx context.Context, content string) (*model.Page, error) {
	page, err := r.DB.CreatePage(content)
	u.LogError(err)
	return page.ToModel(), nil
}

func (r *mutationResolver) DeletePage(ctx context.Context, id string) (bool, error) {
	err := r.DB.DeletePage(id)
	u.LogError(err)
	return err == nil, err
}

func (r *queryResolver) Pages(ctx context.Context) ([]*model.Page, error) {
	modelPages := make([]*model.Page, 0)
	pages, err := r.DB.GetPages()
	if err != nil {
		u.LogError(err)
		return modelPages, err
	}
	for _, page := range pages {
		modelPages = append(modelPages, page.ToModel())
	}
	return modelPages, nil
}

func (r *queryResolver) Page(ctx context.Context, id string) (*model.Page, error) {
	page, err := r.DB.GetPage(id)
	if err != nil {
		u.LogError(err)
		return nil, err
	}
	return page.ToModel(), nil
}
