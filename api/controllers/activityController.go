package controllers

import (
	"net/http"
	"strconv"

	protoActivity "github.com/asciiu/gomo/activity-bulletin/proto"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

// A ResponseActivitySuccess will always contain a status of "successful".
// swagger:model responseActivitySuccess
type ResponseActivitySuccess struct {
	Status string                          `json:"status"`
	Data   *protoActivity.UserActivityPage `json:"data"`
}

// This struct is used in the generated swagger docs,
// and it is not used anywhere.
// swagger:parameters searchActivity
type SearchActivity struct {
	// Required: false
	// In: query
	ObjectID string `json:"objectID"`
	// Required: false
	// In: query
	Page string `json:"page"`
	// Required: false
	// In: query
	PageSize string `json:"pageSize"`
}

type ActivityController struct {
	Bulletin protoActivity.ActivityBulletinClient
}

func NewActivityController(service micro.Service) *ActivityController {
	controller := ActivityController{
		Bulletin: protoActivity.NewActivityBulletinClient("bulletin", service.Client()),
	}

	return &controller
}

// swagger:route GET /protoActivity protoActivity searchActivity
//
// get protoActivity (protected)
//
// Returns a list of protoActivity.
//
// responses:
//  200: responseActivitySuccess "data" will contain array of protoActivity with "status": "success"
func (controller *ActivityController) HandleListActivity(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["jti"].(string)

	objectID := c.QueryParam("objectID")
	pageStr := c.QueryParam("page")
	pageSizeStr := c.QueryParam("pageSize")

	// defaults for page and page size here
	// ignore the errors and assume the values are int
	page, _ := strconv.ParseUint(pageStr, 10, 32)
	pageSize, _ := strconv.ParseUint(pageSizeStr, 10, 32)
	if pageSize == 0 {
		pageSize = 20
	}

	req := protoActivity.ActivityRequest{
		UserID:   userID,
		ObjectID: objectID,
		Page:     uint32(page),
		PageSize: uint32(pageSize),
	}

	r, _ := controller.Bulletin.FindUserActivity(context.Background(), &req)
	if r.Status != "success" {
		response := &ResponseError{
			Status:  r.Status,
			Message: r.Message,
		}

		if r.Status == "fail" {
			return c.JSON(http.StatusBadRequest, response)
		}
		if r.Status == "error" {
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := &ResponseActivitySuccess{
		Status: "success",
		Data:   r.Data,
	}

	return c.JSON(http.StatusOK, response)
}