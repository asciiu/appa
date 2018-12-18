package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	constRes "github.com/asciiu/oldiez/common/constants/response"
	protoOrder "github.com/asciiu/oldiez/order-service/proto/order"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type ResponseOrderSuccess struct {
	Status string `json:"status"`
	Data   *Order `json:"data"`
}

type OrderController struct {
	DB *sql.DB
	// TODO change this to be the OrderClient
	OrderClient protoOrder.OrderService
}

func NewOrderController(db *sql.DB, service micro.Service) *OrderController {
	controller := OrderController{
		DB: db,
		// TODO order client service
		OrderClient: protoOrder.NewOrderService("orders", service.Client()),
	}
	return &controller
}

type NewOrder struct {
	MarketName string  `json:"marketName"`
	Side       string  `json:"side"`
	Size       float64 `json:"size"`
	Type       string  `json:"type"`
}

type Order struct {
	OrderID    string  `json:"orderID"`
	MarketName string  `json:"marketName"`
	Side       string  `json:"side"`
	Size       float64 `json:"size"`
	Type       string  `json:"type"`
	Status     string  `json:"status"`
	CreatedOn  string  `json:"createdOn"`
	UpdatedOn  string  `json:"updatedOn"`
}

func (controller *OrderController) HandlePostOrder(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["jti"].(string)

	newOrder := new(NewOrder)
	err := c.Bind(&newOrder)
	if err != nil {
		response := &ResponseError{
			Status:  constRes.Fail,
			Message: err.Error(),
		}

		return c.JSON(http.StatusBadRequest, response)
	}

	request := protoOrder.NewOrderRequest{
		UserID:     userID,
		MarketName: newOrder.MarketName,
		Side:       newOrder.Side,
		Size:       newOrder.Size,
		Type:       newOrder.Type,
	}

	r, _ := controller.OrderClient.AddOrder(context.Background(), &request)
	if r.Status != constRes.Success {
		res := &ResponseError{
			Status:  r.Status,
			Message: r.Message,
		}

		if r.Status == constRes.Fail {
			return c.JSON(http.StatusBadRequest, res)
		}
		if r.Status == constRes.Error {
			return c.JSON(http.StatusInternalServerError, res)
		}
	}

	order := r.Data.Order
	response := &ResponseOrderSuccess{
		Status: constRes.Success,
		Data: &Order{
			OrderID:    order.OrderID,
			MarketName: order.MarketName,
			Side:       order.Side,
			Size:       order.Size,
			Type:       order.Type,
			Status:     order.Status,
			CreatedOn:  order.CreatedOn,
			UpdatedOn:  order.UpdatedOn,
		},
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

	updateRequest := new(Order)

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
	//userID := claims["jti"].(string)

	request := protoOrder.OrderRequest{
		OrderID: c.Param("orderID"),
	}

	r, _ := controller.OrderClient.FindOrder(context.Background(), &request)
	if r.Status != constRes.Success {
		res := &ResponseError{
			Status:  r.Status,
			Message: r.Message,
		}

		if r.Status == constRes.Fail {
			return c.JSON(http.StatusBadRequest, res)
		}
		if r.Status == constRes.Error {
			return c.JSON(http.StatusInternalServerError, res)
		}
	}

	order := r.Data.Order
	response := &ResponseOrderSuccess{
		Status: constRes.Success,
		Data: &Order{
			OrderID:    order.OrderID,
			MarketName: order.MarketName,
			Side:       order.Side,
			Size:       order.Size,
			Type:       order.Type,
			Status:     order.Status,
			CreatedOn:  order.CreatedOn,
			UpdatedOn:  order.UpdatedOn,
		},
	}

	return c.JSON(http.StatusOK, response)
}

func (controller *OrderController) HandleDeleteOrder(c echo.Context) error {
	//token := c.Get("user").(*jwt.Token)
	//claims := token.Claims.(jwt.MapClaims)
	//userID := claims["jti"].(string)
	request := protoOrder.OrderRequest{
		OrderID: c.Param("orderID"),
	}

	r, _ := controller.OrderClient.CancelOrder(context.Background(), &request)
	if r.Status != constRes.Success {
		res := &ResponseError{
			Status:  r.Status,
			Message: r.Message,
		}

		if r.Status == constRes.Fail {
			return c.JSON(http.StatusBadRequest, res)
		}
		if r.Status == constRes.Error {
			return c.JSON(http.StatusInternalServerError, res)
		}
	}

	response := &ResponseOrderSuccess{
		Status: constRes.Success,
	}

	return c.JSON(http.StatusOK, response)
}
