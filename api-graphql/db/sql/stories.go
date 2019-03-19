package sql

import (
	"database/sql"

	"github.com/asciiu/appa/api-graphql/models"
)

func Unpublish(db *sql.DB, storyID string) error {
	_, err := db.Exec("UPDATE stories SET status = 'unpublish' WHERE id = $1", storyID)
	return err
}

func FindStoryByID(db *sql.DB, storyID string) (*models.Story, error) {
	var s models.Story
	err := db.QueryRow(`SELECT 
	id, 
	author_id, 
	title, 
	content, 
	status FROM stories WHERE id = $1`, storyID).
		Scan(&s.ID,
			&s.AuthorID,
			&s.Title,
			&s.Content,
			&s.Status)

	if err != nil {
		return nil, err
	}
	return &s, nil
}

func StoryTitles(db *sql.DB, status string, page, pageSize uint32) (*models.PagedTitles, error) {
	var total uint32
	queryCount := `SELECT count(*) FROM stories WHERE status = $1`
	err := db.QueryRow(queryCount, status).Scan(&total)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`SELECT 
	   id, 
	   author_id, 
	   title
	   FROM stories where status = $1
	   ORDER BY created_on OFFSET $2 LIMIT $3`, status, page, pageSize)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	stories := make([]*models.Story, 0)
	for rows.Next() {
		s := new(models.Story)
		err := rows.Scan(
			&s.ID,
			&s.AuthorID,
			&s.Title)

		if err != nil {
			return nil, err
		}
		stories = append(stories, s)
	}

	return &models.PagedTitles{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		Stories:  stories,
	}, nil
}

func InsertStory(db *sql.DB, story *models.Story) error {
	sqlStatement := `insert into stories (
		id, 
		author_id, 
		title, 
		content, 
		status) values ($1, $2, $3, $4, $5)`

	_, err := db.Exec(sqlStatement,
		story.ID,
		story.AuthorID,
		story.Title,
		story.Content,
		story.Status)

	return err
}
