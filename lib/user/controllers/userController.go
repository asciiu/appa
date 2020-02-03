package controllers

import (
	"database/sql"
	"fmt"
	"strings"

	queries "github.com/asciiu/appa/lib/user/db/sql"
	"github.com/asciiu/appa/lib/user/models"
)

type UserController struct {
	DB *sql.DB
}

func NewUserController(db *sql.DB) *UserController {
	return &UserController{DB: db}
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (controller *UserController) CreateUser(req *CreateUserRequest) (*models.User, error) {
	user := models.NewUser(req.Username, req.Email, req.Password)
	err := queries.InsertUser(controller.DB, user)

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
		return queries.DeleteUserHard(controller.DB, req.UserID)
	}
	return queries.DeleteUserSoft(controller.DB, req.UserID)
}

func (controller *UserController) GetUser(userID string) (*models.User, error) {
	user, err := queries.FindUserByID(controller.DB, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (controller *UserController) UserEmailVerified(userID string) error {
	return queries.UpdateEmailVerified(controller.DB, userID, true)
}

type UpdateUserRequest struct {
	UserID   string
	Username string
}

func (controller *UserController) UpdateUser(req *UpdateUserRequest) error {
	return queries.UpdateUsername(controller.DB, req.UserID, req.Username)
}
