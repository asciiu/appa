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

func StoryTitles(db *sql.DB, status string, page, pageSize uint32) ([]models.Story, error) {
	//var total uint32
	//queryCount := `SELECT count(*) FROM orders WHERE user_id = $1 AND status like '%' || $2 || '%'`
	//err := db.QueryRow(queryCount, userID, status).Scan(&total)
	//if err != nil {
	//	return nil, err
	//}

	// rows, err := db.Query(`SELECT
	//     id,
	//     user_id,
	//     market_name,
	//     side,
	// 	size,
	// 	price,
	// 	type,
	// 	status,
	// 	created_on,
	// 	updated_on
	// 	FROM orders WHERE user_id = $1 AND status like '%' || $2 || '%'
	// 	OFFSET $3 LIMIT $4`, userID, status, page, pageSize)

	// if err != nil {
	// 	return nil, err
	// }

	// defer rows.Close()

	//orders := make([]*protoOrder.Order, 0)
	//for rows.Next() {
	//	var price sql.NullFloat64
	//	o := new(protoOrder.Order)
	//	err := rows.Scan(
	//		&o.OrderID,
	//		&o.UserID,
	//		&o.MarketName,
	//		&o.Side,
	//		&o.Size,
	//		&price,
	//		&o.Type,
	//		&o.Status,
	//		&o.CreatedOn,
	//		&o.UpdatedOn,
	//	)

	//	if err != nil {
	//		return nil, err
	//	}
	//	if price.Valid {
	//		o.Price = price.Float64
	//	}
	//	orders = append(orders, o)
	//}

	//return &protoOrder.OrdersPage{
	//	Page:     page,
	//	PageSize: pageSize,
	//	Total:    total,
	//	Orders:   orders,
	//}, nil
	return []models.Story{}, nil
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
