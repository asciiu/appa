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
		return []models.User{}, fmt.Errorf("Access denied")
	}

	return []models.User{}, nil
}

func (r *queryResolver) Info(ctx context.Context) (models.User, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return models.User{}, fmt.Errorf("Access denied")
	}

	return *user, nil
}

func (r *queryResolver) FindOrder(ctx context.Context, id string) (*models.Order, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("Access denied")
	}

	return &models.Order{ID: id, Txt: "yeah!"}, nil
}
