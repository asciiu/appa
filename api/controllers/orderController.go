package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	constRes "github.com/asciiu/oldiez/common/constants/response"
	"github.com/labstack/echo"
	micro "github.com/micro/go-micro"
)

type OrderRequest struct {
	CurrencyName string  `json:"currencyName"`
	MarketName   string  `json:"marketName"`
	Side         string  `json:"side"`
	Size         float64 `json:"size"`
}

type OrderController struct {
	DB *sql.DB
	// TODO change this to be the OrderClient
	//UserClient protoUser.UserServiceClient
}

func NewOrderController(db *sql.DB, service micro.Service) *OrderController {
	controller := OrderController{
		DB: db,
		// TODO order client service
		//Orderlient: protoOrder.NewOrderServiceClient("orders", service.Client()),
	}
	return &controller
}

// swagger:route POST /orders orders order
//
// creates a new order in the system (open)
//
// blah blah blah
//
// responses:
//  200: responseSuccess "data" will be non null with "status": "success"
func (controller *OrderController) HandlePostOrder(c echo.Context) error {
	//token := c.Get("user").(*jwt.Token)
	//claims := token.Claims.(jwt.MapClaims)
	//userID := claims["jti"].(string)

	var orderRequest OrderRequest
	err := json.NewDecoder(c.Request().Body).Decode(&orderRequest)
	if err != nil {
		response := &ResponseError{
			Status:  constRes.Fail,
			Message: err.Error(),
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
// updates an order (protected)
//
// todo
//
// responses:
func (controller *OrderController) HandleUpdateOrder(c echo.Context) error {
	//token := c.Get("user").(*jwt.Token)
	//claims := token.Claims.(jwt.MapClaims)
	//orderID := c.Param("orderId")
	//userID := claims["jti"].(string)

	updateRequest := new(OrderRequest)

	err := json.NewDecoder(c.Request().Body).Decode(&updateRequest)
	if err != nil {
		response := &ResponseError{
			Status:  constRes.Fail,
			Message: err.Error(),
		}

		return c.JSON(http.StatusBadRequest, response)
	}

	response := &ResponseSuccess{
		Status: constRes.Success,
	}

	return c.JSON(http.StatusOK, response)
}

func (controller *OrderController) HandleGetOrders(c echo.Context) error {
	//token := c.Get("user").(*jwt.Token)
	//claims := token.Claims.(jwt.MapClaims)
	//orderID := c.Param("orderId")
	//userID := claims["jti"].(string)

	response := &ResponseSuccess{
		Status: constRes.Success,
	}

	return c.JSON(http.StatusOK, response)
}

func (controller *OrderController) HandleGetOrder(c echo.Context) error {
	//token := c.Get("user").(*jwt.Token)
	//claims := token.Claims.(jwt.MapClaims)
	//orderID := c.Param("orderId")
	//userID := claims["jti"].(string)

	response := &ResponseSuccess{
		Status: constRes.Success,
	}

	return c.JSON(http.StatusOK, response)
}

func (controller *OrderController) HandleDeleteOrder(c echo.Context) error {
	//token := c.Get("user").(*jwt.Token)
	//claims := token.Claims.(jwt.MapClaims)
	//orderID := c.Param("orderId")
	//userID := claims["jti"].(string)

	response := &ResponseSuccess{
		Status: constRes.Success,
	}

	return c.JSON(http.StatusOK, response)
}
