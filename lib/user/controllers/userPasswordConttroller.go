package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	queries "github.com/asciiu/appa/lib/user/db/sql"
	"github.com/asciiu/appa/lib/user/models"
	"golang.org/x/crypto/bcrypt"
)

type UserPasswordController struct {
	DB *sql.DB
}

func NewUserPasswordController(db *sql.DB) *UserPasswordController {
	return &UserPasswordController{DB: db}
}

type ChangePasswordRequest struct {
	UserID      string `json:"userID"`
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"oldPassword"`
}

// ChangePassword - Changes the user's password. Password is updated when the request's
// old password matches the current user's password hash.
func (controller *UserPasswordController) ChangePassword(req *ChangePasswordRequest) error {
	user, err := queries.FindUserByID(controller.DB, req.UserID)

	switch {
	case err == nil:
		if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.OldPassword)) == nil {

			err := queries.UpdatePassword(controller.DB, req.UserID, models.HashAndSalt([]byte(req.NewPassword)))
			if err != nil {
				return err
			}
		} else {
			return errors.New("old password incorrect")
		}

	case strings.Contains(err.Error(), "no rows in result set"):
		return fmt.Errorf("user ID not found: %s", req.UserID)

	default:
		return err
	}

	return nil
}
