package gopg

import (
	"github.com/asciiu/appa/lib/user/models"
	"github.com/go-pg/pg/v10"
)

func DeleteUserHard(db *pg.DB, userID string) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", userID)
	return err
}

func FindUserByEmail(db *pg.DB, email string) (*models.User, error) {
	u := new(models.User)

	err := db.Model(u).Where("email = ?", email).Select()

	return u, err
}

func FindUserByID(db *pg.DB, userID string) (*models.User, error) {
	u := new(models.User)

	err := db.Model(u).Where("id = ?", userID).Select()

	return u, err
}

func InsertUser(db *pg.DB, user *models.User) error {
	return db.Insert(user)
}

func UpdateUser(db *pg.DB, user *models.User) error {
	return db.Update(user)
}
