package graphql

import (
	"database/sql"
)

type Resolver struct {
	DB *sql.DB
	//StoryClient protoStory.StoryService
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

// type mutationResolver struct{ *Resolver }

// func (r *mutationResolver) Signup(ctx context.Context, email string, username string, password string) (*models.User, error) {
// 	panic("not implemented")
// }
// func (r *mutationResolver) Login(ctx context.Context, email string, password string, remember bool) (*Token, error) {
// 	panic("not implemented")
// }
// func (r *mutationResolver) CreateStory(ctx context.Context, title string, json string) (string, error) {
// 	panic("not implemented")
// }
// func (r *mutationResolver) UpdateStory(ctx context.Context, id string, title string, json string, status string) (bool, error) {
// 	panic("not implemented")
// }

// type queryResolver struct{ *Resolver }

// func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
// 	panic("not implemented")
// }
// func (r *queryResolver) Balances(ctx context.Context) ([]*models.Balance, error) {
// 	panic("not implemented")
// }
// func (r *queryResolver) Info(ctx context.Context) (*models.User, error) {
// 	panic("not implemented")
// }
// func (r *queryResolver) GetUser(ctx context.Context) (*models.User, error) {
// 	panic("not implemented")
// }
// func (r *queryResolver) UserSummary(ctx context.Context) (*models.UserSummary, error) {
// 	panic("not implemented")
// }
// func (r *queryResolver) FindOrder(ctx context.Context, id string) (*models.Order, error) {
// 	panic("not implemented")
// }
// func (r *queryResolver) ListStories(ctx context.Context) ([]*models.Story, error) {
// 	panic("not implemented")
// }
