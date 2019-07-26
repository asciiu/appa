package sql

import (
	"database/sql"
)

func InsertWriter(db *sql.DB, userID, title string) error {
	sqlStatement := `insert into writers (
		user_id, 
		title) values ($1, $2)`
	_, err := db.Exec(sqlStatement,
		userID,
		title,
	)

	return err
}
