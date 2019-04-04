package apiql

import (
	"context"
	"fmt"

	"github.com/asciiu/appa/api-graphql/auth"
	repo "github.com/asciiu/appa/api-graphql/db/sql"
	"github.com/asciiu/appa/api-graphql/models"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Balances(ctx context.Context) ([]*models.Balance, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return []*models.Balance{}, fmt.Errorf("unauthorized")
	}

	balances, err := repo.FindUserBalances(r.DB, user.ID)
	return balances, err
}

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

func (r *queryResolver) UserSummary(ctx context.Context) (*models.UserSummary, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("unauthorized")
	}
	summary := models.UserSummary{
		User: user,
	}

	return &summary, nil
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
