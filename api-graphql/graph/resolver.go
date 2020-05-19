package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"database/sql"

	protoStory "github.com/asciiu/appa/story-service/proto/story"
)

type Resolver struct {
	DB          *sql.DB
	StoryClient protoStory.StoryService
}
