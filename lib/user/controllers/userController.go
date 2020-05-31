package controllers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/asciiu/appa/lib/user/models"
)

type UserController struct {
	userRepo models.UserRepo
}

func NewUserController(userRepo models.UserRepo) *UserController {
	return &UserController{userRepo: userRepo}
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (controller *UserController) CreateUser(username, email, password string) (*models.User, error) {
	if username == "" || email == "" || password == "" {
		return nil, errors.New("cannot create user with empty value")
	}
	user := models.NewUser(username, email, password)
	err := controller.userRepo.InsertUser(user)

	switch {
	case err == nil:
		return user, nil

	case strings.Contains(err.Error(), "violates unique constraint \"users_email_key\""):
		return nil, fmt.Errorf("email already exists")

	default:
		return nil, err
	}
}

type DeleteUserRequest struct {
	UserID string
	IsHard bool
}

func (controller *UserController) DeleteUser(req *DeleteUserRequest) error {
	if req.IsHard {
		//return queries.DeleteUserHard(controller.DB, req.UserID)
		return controller.userRepo.DeleteUserHard(req.UserID)
	}
	return controller.userRepo.DeleteUserSoft(req.UserID)
}

func (controller *UserController) GetUser(userID string) (*models.User, error) {
	user, err := controller.userRepo.FindUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (controller *UserController) FindUserByEmail(email string) (*models.User, error) {
	return controller.userRepo.FindUserByEmail(email)
}

func (controller *UserController) FindUserByID(userID string) (*models.User, error) {
	return controller.userRepo.FindUserByID(userID)
}

func (controller *UserController) UserEmailVerified(userID string) error {
	_, err := controller.userRepo.UpdateEmailVerified(userID, true)
	return err
}

type UpdateUserRequest struct {
	UserID   string
	Username string
}

func (controller *UserController) UpdateUser(req *UpdateUserRequest) error {
	_, err := controller.userRepo.UpdateUsername(req.UserID, req.Username)
	return err
}
