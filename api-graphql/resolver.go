package apiql

import (
	"database/sql"
)

type Resolver struct {
	DB *sql.DB
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
