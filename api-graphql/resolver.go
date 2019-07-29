package apiql

import (
	"database/sql"

	protoStory "github.com/asciiu/appa/story-service/proto/story"
)

type Resolver struct {
	DB          *sql.DB
	StoryClient protoStory.StoryService
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
