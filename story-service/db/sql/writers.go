package sql

import (
	"database/sql"
)

// func DeleteOrder(db *sql.DB, orderID, userID string) error {
// 	_, err := db.Exec("DELETE FROM orders WHERE id = $1 and user_id = $2", orderID, userID)
// 	return err
// }

func InsertWriter(db *sql.DB, userID, title string) error {
	sqlStatement := `insert into authors (
		user_id, 
		title) values ($1, $2)`
	_, err := db.Exec(sqlStatement,
		userID,
		title,
	)

	return err
}
