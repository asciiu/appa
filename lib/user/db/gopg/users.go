package gopg

import (
	"github.com/asciiu/appa/lib/user/models"
	"github.com/go-pg/pg/v10"
)

type UserRepo struct {
	db *pg.DB
}

func NewUserRepo(db *pg.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) DeleteUserHard(userID string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", userID)
	return err
}

func (r *UserRepo) DeleteUserSoft(userID string) error {
	_, err := r.db.Exec("UPDATE users SET deleted_on = now() WHERE id = $1", userID)
	return err
}

func (r *UserRepo) FindUserByEmail(email string) (*models.User, error) {
	u := new(models.User)

	err := r.db.Model(u).Where("email = ?", email).Select()

	return u, err
}

func (r *UserRepo) FindUserByID(userID string) (*models.User, error) {
	u := new(models.User)

	err := r.db.Model(u).Where("id = ?", userID).Select()

	return u, err
}

func (r *UserRepo) InsertUser(user *models.User) error {
	return r.db.Insert(user)
}

func (r *UserRepo) UpdateUser(user *models.User) error {
	return r.db.Update(user)
}

func (r *UserRepo) UpdatePassword(userID, hash string) (*models.User, error) {
	user := new(models.User)
	_, err := r.db.Model(user).Set("password_hash = ?", hash).Where("id = ?", userID).Update()
	return user, err
}

func (r *UserRepo) UpdateUsername(userID, username string) (*models.User, error) {
	user := new(models.User)
	_, err := r.db.Model(user).Set("username = ?", username).Where("id = ?", userID).Update()
	return user, err
}

func (r *UserRepo) UpdateEmailVerified(userID string, verified bool) (*models.User, error) {
	user := new(models.User)
	_, err := r.db.Model(user).Set("email_verified = ?", verified).Where("id = ?", userID).Update()
	return user, err
}
