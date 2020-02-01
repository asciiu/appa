package handlers

import (
	"database/sql"
	"net/http"

	constRes "github.com/asciiu/appa/lib/constants/response"
	userControllers "github.com/asciiu/appa/lib/user/controllers"
	userModels "github.com/asciiu/appa/lib/user/models"
	"github.com/labstack/echo"
	micro "github.com/micro/go-micro"
)

type SessionHandler struct {
	DB             *sql.DB
	UserController *userControllers.UserController
}

func NewSessionHandler(db *sql.DB, service micro.Service) *SessionHandler {
	handler := SessionHandler{
		DB:             db,
		UserController: userControllers.NewUserController(db),
	}
	return &handler
}

type UserData struct {
	User *userModels.User `json:"user"`
}

// A ResponseSessionSuccess will always contain a status of "successful".
// swagger:model ResponseSessionSuccess
type ResponseSessionSuccess struct {
	Status string    `json:"status"`
	Data   *UserData `json:"data"`
}

// swagger:route GET /session session sessionBegin
//
// create a new session for a user (protected)
//
// Creates a new session for an authenticated user. The session data will eventually contain
// whatever info you need to begin a new session. At the moment the response data mirrors
// login data. This endpoint depends on the user-service. If the user-service
// is unreachable, a 410 with a status of "error" will be returned.
//
// responses:
//  200: ResponseSessionSuccess data will be non null with status "success"
//  410: responseError the user-service is unreachable with status "error"
func (handler *SessionHandler) HandleSession(c echo.Context) error {
	ctx := c.(*UserContext)
	usr := ctx.User

	response := ResponseSessionSuccess{
		Status: constRes.Success,
		Data:   &UserData{User: user},
	}

	return c.JSON(http.StatusOK, &response)
}
