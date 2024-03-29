package sql

import (
	"database/sql"

	"github.com/asciiu/appa/lib/user/models"
)

func DeleteUserHard(db *sql.DB, userID string) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", userID)
	return err
}

func DeleteUserSoft(db *sql.DB, userID string) error {
	_, err := db.Exec("UPDATE users SET deleted_on = now() WHERE id = $1", userID)
	return err
}

func FindUserByEmail(db *sql.DB, email string) (*models.User, error) {
	var u models.User
	var avatarURL sql.NullString
	err := db.QueryRow(`SELECT 
	id, 
	username, 
	avatar_url,
	email, 
	email_verified, 
	password_hash FROM users WHERE email = $1`, email).
		Scan(&u.ID,
			&u.Username,
			&avatarURL,
			&u.Email,
			&u.EmailVerified,
			&u.PasswordHash)

	if err != nil {
		return nil, err
	}
	if avatarURL.Valid {
		u.AvatarURL = avatarURL.String
	}
	return &u, nil
}

func FindUserByID(db *sql.DB, userID string) (*models.User, error) {
	var u models.User
	var avatarURL sql.NullString
	err := db.QueryRow(`SELECT 
	id, 
	username, 
	avatar_url,
	email, 
	email_verified, 
	password_hash FROM users WHERE id = $1`, userID).
		Scan(&u.ID,
			&u.Username,
			&avatarURL,
			&u.Email,
			&u.EmailVerified,
			&u.PasswordHash)

	if err != nil {
		return nil, err
	}
	if avatarURL.Valid {
		u.AvatarURL = avatarURL.String
	}
	return &u, nil
}

func InsertUser(db *sql.DB, user *models.User) error {
	sqlStatement := `insert into users (
		id, 
		email,
		username, 
		email_verified, 
		password_hash) values ($1, $2, $3, $4, $5)`

	_, err := db.Exec(sqlStatement,
		user.ID,
		user.Email,
		user.Username,
		user.EmailVerified,
		user.PasswordHash)

	return err
}

func UpdatePassword(db *sql.DB, userID, hash string) error {
	_, err := db.Exec("UPDATE users SET password_hash = $1 WHERE id = $2",
		hash, userID)

	return err
}

func UpdateUsername(db *sql.DB, userID, username string) error {
	_, err := db.Exec("UPDATE users SET username = $1 WHERE id = $2",
		username, userID)

	return err
}

func UpdateEmailVerified(db *sql.DB, userID string, emailVerified bool) error {
	_, err := db.Exec("UPDATE users SET email_verified = $1 WHERE id = $2",
		emailVerified, userID)

	return err
}
