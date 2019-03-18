package apiql

import (
	"context"
	"fmt"

	"github.com/asciiu/appa/apiql/auth"
	"github.com/asciiu/appa/apiql/models"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]models.User, error) {
	if user := auth.ForContext(ctx); user == nil {
		return []models.User{}, fmt.Errorf("unauthorized")
	}

	return []models.User{}, nil
}

func (r *queryResolver) Info(ctx context.Context) (models.User, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return models.User{}, fmt.Errorf("unauthorized")
	}

	return *user, nil
}

func (r *queryResolver) GetUser(ctx context.Context) (*models.User, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("unauthorized")
	}

	return user, nil
}

func (r *queryResolver) FindOrder(ctx context.Context, id string) (*models.Order, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("unauthorized")
	}

	return &models.Order{ID: id, Txt: "yeah!"}, nil
}

func (r *queryResolver) ListStories(ctx context.Context) ([]models.Story, error) {
	//user := auth.ForContext(ctx)
	//if user == nil {
	//	return []models.Story{}, fmt.Errorf("unauthorized")
	//}

	stories := []models.Story{}
	return stories, nil
}
