package gopg

import (
	"github.com/asciiu/appa/lib/user/models"
	"github.com/go-pg/pg/v10"
)

func FindUserByEmail(db *pg.DB, email string) (*models.User, error) {
	u := new(models.User)

	err := db.Model(u).Where("email = ?", email).Select()

	return u, err
}
