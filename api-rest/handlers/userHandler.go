package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	constRes "github.com/asciiu/appa/lib/constants/response"
	userControllers "github.com/asciiu/appa/lib/user/controllers"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/net/context"
)

type UseHandler struct {
	DB                 *sql.DB
	PasswordController *userControllers.UserPasswordController
	UserController *userControllers.UserController
}

// swagger:parameters UpdateUser
type UpdateUserRequest struct {
	// Optional.
	// in: body
	First string `json:"first"`
	// Optional.
	// in: body
	Last string `json:"last"`
	// Optional. Note: we need to validate these!
	// in: body
	Email string `json:"email"`
}

func NewUserController(db *sql.DB) *UseHandler {
	handler := UseHandler{
		DB:                 db,
		PasswordController: userControllers.NewUserPasswordController(db),
		UserController: userControllers.NewUserController(db),
	}
	return &handler
}

// swagger:route PUT /users/:id/changepassword users ChangePassword
//
// change a user's password (protected)
//
// Allows an authenticated user to change their password. The url param is the user's id.
//
// responses:
//  200: responseSuccess the status will be "success" with data null.
//  400: responseError you did something wrong here with status "fail". Hopefully, the message is descriptive enough.
//  401: responseError the user ID in url param does not match with status "fail".
//  410: responseError the user-service is unreachable with status "error"
//  500: responseError the message will state what the internal server error was with "status": "error"
func (handler *UseHandler) HandleChangePassword(c echo.Context) error {
	ctx := c.(*UserContext)
	user := ctx.User
	paramID := c.Param("userID")
	userID := user.ID

	if paramID != userID {
		response := &ResponseError{
			Status:   constRes.Fail,
			Messages: []string{"unauthorized"},
		}

		return c.JSON(http.StatusUnauthorized, response)
	}

	passwordRequest := new(userControllers.ChangePasswordRequest)
	err := c.Bind(passwordRequest)
	if err != nil {
		response := &ResponseError{
			Status:   constRes.Fail,
			Messages: []string{err.Error()},
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	if err = c.Validate(passwordRequest); err != nil {
		msgs := strings.Split(err.Error(), "\n")
		response := &ResponseError{
			Status:   constRes.Fail,
			Messages: msgs,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := handler.PasswordController.ChangePassword(passwordRequest); err != nil {
		response := &ResponseError{
			Status:   constRes.Fail,
			Messages: []string{err.Error()},
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	response := &ResponseSuccess{
		Status: constRes.Success,
	}

	return c.JSON(http.StatusOK, response)
}

// swagger:route PUT /users/:id users UpdateUser
//
// updates user info (protected)
//
// You can change the user's first, last, or email. Note we need to implement a secure method of
// verifing the user's new email. This has yet to be implemented.
//
// responses:
//  200: responseSuccess "data" will contain updated user data with "status": "success"
//  400: responseError message in badrequest should be descriptive with "status": "fail"
//  401: responseError unauthorized user because of incorrect url param with "status": "fail"
//  410: responseError the user-service is unreachable with status "error"
//  500: responseError the message will state what the internal server error was with "status": "error"
func (handler *UseHandler) HandleUpdateUser(c echo.Context) error {
	ctx := c.(*UserContext)
	user := ctx.User
	paramID := c.Param("userID")
	userID := user.ID

	if paramID != userID {
		response := &ResponseError{
			Status:   constRes.Fail,
			Messages: []string{"unauthorized"},
		}

		return c.JSON(http.StatusUnauthorized, response)
	}

	updateRequest := new(userControllers.UpdateUserRequest)
	err := c.Bind(updateRequest)
	if err != nil {
		response := &ResponseError{
			Status:   constRes.Fail,
			Messages: []string{err.Error()},
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(updateRequest); err != nil {
		msgs := strings.Split(err.Error(), "\n")
		response := &ResponseError{
			Status:   constRes.Fail,
			Messages: msgs,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

    if err := handler.UserController.UpdateUser(updateRequest); err != nil {
		msgs := strings.Split(err.Error(), "\n")
		response := &ResponseError{
			Status:   constRes.Fail,
			Messages: msgs,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	user.Username = updateRequest.Username
	response := ResponseSessionSuccess{
		Status: constRes.Success,
		Data:   &UserData{User: user},
	}

	return c.JSON(http.StatusOK, response)
}

// swagger:route POST /signup authentication signup
//
// user registration (open)
//
// Registers a new user. Expects email to be unique. Duplicate email will result
// in a bad request.
//
// responses:
//  200: responseSuccess "data" will be non null with "status": "success"
//  400: responseError message should relay information with regard to bad request with "status": "fail"
//  410: responseError the user-service is not reachable. The user-service is a microservice that runs independantly from the api. When we take it offline you will receive this error.
//  500: responseError the message will state what the internal server error was with "status": "error"
func (handler *UseHandler) HandleSignup(c echo.Context) error {

	 signupRequest := new(userControllers.CreateUserRequest)
	 err := c.Bind(signupRequest)
	 if err != nil {
		 response := &ResponseError{
			 Status:   constRes.Fail,
			 Messages: []string{err.Error()},
		 }
		 return c.JSON(http.StatusBadRequest, response)
	 }

	 if err := c.Validate(signupRequest); err != nil {
		msgs := strings.Split(err.Error(), "\n")
		response := &ResponseError{
			Status:   constRes.Fail,
			Messages: msgs,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	user, err := handler.UserController.CreateUser(signupRequest)
	if err != nil {
		msgs := strings.Split(err.Error(), "\n")
		response := &ResponseError{
			Status:   constRes.Fail,
			Messages: msgs,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	response := ResponseSessionSuccess{
		Status: constRes.Success,
		Data:   &UserData{User: user},
	}

	return c.JSON(http.StatusOK, response)
}