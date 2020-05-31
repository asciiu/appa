package controllers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/asciiu/appa/lib/user/models"
	"golang.org/x/crypto/bcrypt"
)

type UserPasswordController struct {
	userRepo models.UserRepo
}

func NewUserPasswordController(userRepo models.UserRepo) *UserPasswordController {
	return &UserPasswordController{userRepo: userRepo}
}

type ChangePasswordRequest struct {
	UserID      string `json:"userID"`
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"oldPassword"`
}

// ChangePassword - Changes the user's password. Password is updated when the request's
// old password matches the current user's password hash.
func (controller *UserPasswordController) ChangePassword(req *ChangePasswordRequest) error {
	user, err := controller.userRepo.FindUserByID(req.UserID)

	switch {
	case err == nil:
		if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.OldPassword)) == nil {

			_, err := controller.userRepo.UpdatePassword(req.UserID, models.HashAndSalt([]byte(req.NewPassword)))
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
