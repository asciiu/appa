package sql

import (
	"database/sql"

	"github.com/asciiu/appa/user-service/models"
)

// DeleteUserHard - delete user from DB
func DeleteUserHard(db *sql.DB, userID string) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", userID)
	return err
}

// DeleteUserSoft - sets the deleated_on column
func DeleteUserSoft(db *sql.DB, userID string) error {
	_, err := db.Exec("UPDATE users SET deleted_on = now() WHERE id = $1", userID)
	return err
}

// FindUserByEmail - find a user by email address
func FindUserByEmail(db *sql.DB, email string) (*models.User, error) {
	var u models.User
	err := db.QueryRow(`SELECT
	    id,
	    username,
	    email,
	    email_verified,
	    password_hash,
	    created_on,
	    updated_on
	    FROM users WHERE email = $1`, email).
		Scan(&u.ID,
			&u.Username,
			&u.Email,
			&u.EmailVerified,
			&u.PasswordHash,
			&u.CreatedOn,
			&u.UpdatedOn)

	if err != nil {
		return nil, err
	}
	return &u, nil
}

// FindUserByID - find users by the user id
func FindUserByID(db *sql.DB, userID string) (*models.User, error) {
	var u models.User
	err := db.QueryRow(`SELECT 
	    id, 
	    username, 
	    email, 
	    email_verified, 
	    password_hash,
	    created_on,
		updated_on 
		FROM users WHERE id = $1`, userID).
		Scan(&u.ID,
			&u.Username,
			&u.Email,
			&u.EmailVerified,
			&u.PasswordHash,
			&u.CreatedOn,
			&u.UpdatedOn)

	if err != nil {
		return nil, err
	}
	return &u, nil
}

// InsertUser - insert a new user
func InsertUser(db *sql.DB, user *models.User) (*models.User, error) {
	sqlStatement := `insert into users (
		id, 
		username, 
		email, 
		email_verified, 
		password_hash,
		created_on,
		updated_on) values ($1, $2, $3, $4, $5, $6, $7)`
	_, err := db.Exec(sqlStatement,
		user.ID,
		user.Username,
		user.Email,
		user.EmailVerified,
		user.PasswordHash,
		user.CreatedOn,
		user.UpdatedOn,
	)

	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUserPassword - update a user's password
func UpdateUserPassword(db *sql.DB, userID, hash string) error {
	_, err := db.Exec("UPDATE users SET password_hash = $1 WHERE id = $2", hash, userID)
	return err
}

// func UpdateUserInfo(db *sql.DB, user *models.User) (*models.User, error) {
// 	sqlStatement := `UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4`
// 	_, err := db.Exec(sqlStatement, user.First, user.Last, user.Email, user.ID)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }
