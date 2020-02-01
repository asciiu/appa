package handlers

import (
	userModels "github.com/asciiu/appa/lib/user/models"
	"github.com/labstack/echo/v4"
)

type UserContext struct {
	echo.Context
	User *userModels.User
}

// A ResponseSuccess will always contain a status of "successful".
// This response may or may not include data encapsulating the user information.
// swagger:model responseError
type ResponseError struct {
	Status   string   `json:"status"`
	Messages []string `json:"messages"`
}
