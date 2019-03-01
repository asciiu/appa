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
