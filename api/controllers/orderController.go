package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	orders "github.com/asciiu/gomo/order-service/proto/order"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	"golang.org/x/net/context"
)

type OrderController struct {
	DB     *sql.DB
	Orders orders.OrderServiceClient
}

type OrderTemp struct {
	orderID            string
	apiKeyID           string //Key id used for the order? Remember why we have this?
	exchangeOrderID    string
	baseCurrency       string // "BTC",
	baseCurrencyLong   string // "Bitcoin", //As above
	marketCurrency     string // "LTC",
	marketCurrencyLong string // "Litecoin", //Only bittrex seems to have this, pass the short name if doesn't exist
	minTradeSize       string //"0.001", //string
	marketName         string // "LTCBTC", //Convention is market+base this is our name
	//marketPrice: "0.41231231", //String Last price from socket for the pair in the exchange
	//?btcPrice: "0.41231231", //String This is a shortcut for me not to calculate we can discuss it
	//?fiatPrice: "1.35",  //Stting This is a shortcut for me not to calculate we can discuss it
	exchange           string // "binance"
	exchangeMarketName string // "LTC-BTC", //Some exchanges put dash others reverse them i.e. BTCLTC,
	orderType          string // limit, market, stop, fake_market, see above.
	rate               string //String
	baseQuantity       float64
	quantity           float64 // baseQuantity / rate
	quantityRemaining  float64 // how many
	side               string  // buy, sell
	conditions         string
	status             string //open, draft, closed,
	createdAt          int64  //integer
}

// swagger:parameters addOrder
type OrderRequest struct {
	// Required.
	// in: body
	KeyID string `json:"keyID"`
	// Required.
	// in: body
	MarketName string `json:"marketName"`
	// Required.
	// in: body
	Side string `json:"side"`
	// Required. Valid order types are "BuyOrder", "SellOrder", "PaperBuyOrder", "PaperSellOrder". Order not of these types will be ignored.
	// in: body
	OrderType string `json:"orderType"`
	// Required for buy side when order is first in chain
	// in: body
	BaseQuantity float64 `json:"baseQuantity"`
	// Required for buy side on chained orders
	// in: body
	BasePercent float64 `json:"basePercent"`
	// Required for sell side when an order is first in a chain
	// in: body
	CurrencyQuantity float64 `json:"currencyQuantity"`
	// Required for sell side for all orders that are chained
	// in: body
	CurrencyPercent float64 `json:"currencyPercent"`
	// Required.
	// in: body
	Conditions string `json:"conditions"`

	// Optional parent order ID to add this chain of orders to
	ParentOrderID string `json:"parentOrderID"`
}

// swagger:parameters updateOrder
type UpdateOrderRequest struct {
	// Optional.
	// in: body
	OrderType string `json:"orderType"`
	// Optional.
	// in: body
	Price float64 `json:"price"`
	// Optional.
	// in: body
	BaseQuantity float64 `json:"baseQuantity"`
	// Optional.
	// in: body
	Conditions string `json:"conditions"`
}

// A ResponseKeySuccess will always contain a status of "successful".
// swagger:model responseOrderSuccess
type ResponseOrderSuccess struct {
	Status string                `json:"status"`
	Data   *orders.UserOrderData `json:"data"`
}

// A ResponseKeysSuccess will always contain a status of "successful".
// swagger:model responseOrdersSuccess
type ResponseOrdersSuccess struct {
	Status string                 `json:"status"`
	Data   *orders.UserOrdersData `json:"data"`
}

func NewOrderController(db *sql.DB) *OrderController {
	// Create a new service. Optionally include some options here.
	service := k8s.NewService(micro.Name("apikey.client"))
	service.Init()

	controller := OrderController{
		DB:     db,
		Orders: orders.NewOrderServiceClient("orders", service.Client()),
	}
	return &controller
}

// swagger:route GET /orders/:orderID orders getOrder
//
// show order (protected)
//
// Get info about an order.
//
// responses:
//  200: responseOrderSuccess "data" will contain order stuffs with "status": "success"
//  500: responseError the message will state what the internal server error was with "status": "error"
func (controller *OrderController) HandleGetOrder(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["jti"].(string)
	orderID := c.Param("orderID")

	getRequest := orders.GetUserOrderRequest{
		OrderID: orderID,
		UserID:  userID,
	}

	r, _ := controller.Orders.GetUserOrder(context.Background(), &getRequest)
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

	response := &ResponseOrderSuccess{
		Status: "success",
		Data:   r.Data,
	}

	return c.JSON(http.StatusOK, response)
}

// swagger:route GET /orders orders getAllOrders
//
// get all orders (protected)
//
// Currently returns all orders. Eventually going to add params to filter orders.
//
// responses:
//  200: responseOrdersSuccess "data" will contain a list of order info with "status": "success"
//  500: responseError the message will state what the internal server error was with "status": "error"
func (controller *OrderController) HandleListOrders(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["jti"].(string)

	getRequest := orders.GetUserOrdersRequest{
		UserID: userID,
	}

	r, _ := controller.Orders.GetUserOrders(context.Background(), &getRequest)
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

	response := &ResponseOrdersSuccess{
		Status: "success",
		Data:   r.Data,
	}

	return c.JSON(http.StatusOK, response)
}

func fail(c echo.Context, msg string) error {
	response := &ResponseError{
		Status:  "fail",
		Message: msg,
	}

	return c.JSON(http.StatusBadRequest, response)
}

// swagger:route POST /orders orders addOrder
//
// create a new order  (protected)
//
// This will create a new order in the system.
//
// responses:
//  200: responseOrdersSuccess "data" will contain list of orders with "status": "success"
//  400: responseError missing params with "status": "fail"
//  500: responseError the message will state what the internal server error was with "status": "error"
func (controller *OrderController) HandlePostOrder(c echo.Context) error {
	defer c.Request().Body.Close()
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["jti"].(string)

	ordrs := make([]*OrderRequest, 0)
	requests := make([]*orders.OrderRequest, 0)
	dec := json.NewDecoder(c.Request().Body)

	_, err := dec.Token()
	if err != nil {
		return fail(c, err.Error())
	}

	// read all orders from array
	for dec.More() {
		var o OrderRequest

		if err := dec.Decode(&o); err != nil {
			return fail(c, "expected an array")
		}
		ordrs = append(ordrs, &o)
	}

	// error check all orders
	for i, order := range ordrs {
		// side, market name, and api key are required
		if order.Side == "" || order.MarketName == "" || order.KeyID == "" {
			return fail(c, "side, marketName, and apiKeyID required!")
		}

		// assume the first order is head of a chain if the ParentOrderID is empty
		// this means that a new chain of orders has been submitted because the
		// ParentOrderID has not been assigned yet.
		if i == 0 && order.ParentOrderID == "" && order.Side == "buy" && order.BasePercent == 0.0 {
			return fail(c, "head buy in chain requires a basePercent")
		}

		// if the head order side is sell we need a currency quantity
		if i == 0 && order.ParentOrderID == "" && order.Side == "sell" && order.CurrencyPercent == 0.0 {
			return fail(c, "head sell in chain requires a currencyPercent")
		}

		// need to use basePercent for chained buys
		//if i != 0 && order.Side == "buy" && order.BasePercent == 0.0 {
		//	return fail(c, "chained buys require a basePercent")
		//}

		//// need to use currencyPercent for chained buys
		//if i != 0 && order.Side == "sell" && order.CurrencyQuantity == 0.0 {
		//	return fail(c, "chained sells require a currencyPercent")
		//}

		// market name should be formatted as
		// currency-base (e.g. ADA-BTC)
		if !strings.Contains(order.MarketName, "-") {
			return fail(c, "marketName must be currency-base: e.g. ADA-BTC")
		}

		if order.ParentOrderID == "" {
			order.ParentOrderID = "00000000-0000-0000-0000-000000000000"
		}

		request := orders.OrderRequest{
			UserID:           userID,
			KeyID:            order.KeyID,
			MarketName:       order.MarketName,
			Side:             order.Side,
			Conditions:       order.Conditions,
			OrderType:        order.OrderType,
			BaseQuantity:     order.BaseQuantity,
			BasePercent:      order.BasePercent,
			CurrencyQuantity: order.CurrencyQuantity,
			CurrencyPercent:  order.CurrencyPercent,
			ParentOrderID:    order.ParentOrderID,
		}
		requests = append(requests, &request)
	}

	orderRequests := orders.OrdersRequest{
		Orders: requests,
	}

	// add order returns nil for error
	r, _ := controller.Orders.AddOrders(context.Background(), &orderRequests)
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

	response := &ResponseOrdersSuccess{
		Status: "success",
		Data:   r.Data,
	}

	return c.JSON(http.StatusOK, response)
}

// swagger:route PUT /orders/:orderID orders updateOrder
//
// update and order (protected)
//
// You can only update pending orders.
//
// responses:
//  200: responseOrderSuccess "data" will contain order info with "status": "success"
//  400: responseError missing params with "status": "fail"
//  500: responseError the message will state what the internal server error was with "status": "error"
func (controller *OrderController) HandleUpdateOrder(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["jti"].(string)
	orderID := c.Param("orderID")

	orderRequest := UpdateOrderRequest{}

	err := json.NewDecoder(c.Request().Body).Decode(&orderRequest)
	if err != nil {
		response := &ResponseError{
			Status:  "fail",
			Message: err.Error(),
		}

		return c.JSON(http.StatusBadRequest, response)
	}

	// client can only update description
	updateRequest := orders.OrderRequest{
		OrderID:      orderID,
		UserID:       userID,
		Conditions:   orderRequest.Conditions,
		BaseQuantity: orderRequest.BaseQuantity,
	}

	r, _ := controller.Orders.UpdateOrder(context.Background(), &updateRequest)
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

	response := &ResponseOrderSuccess{
		Status: "success",
		Data:   r.Data,
	}

	return c.JSON(http.StatusOK, response)
}

// swagger:route DELETE /orders/:orderID orders deleteOrder
//
// Remove and order (protected)
//
// Cannot remove orders that have already executed.
//
// responses:
//  200: responseOrderSuccess data will be null with "status": "success"
//  500: responseError the message will state what the internal server error was with "status": "error"
func (controller *OrderController) HandleDeleteOrder(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["jti"].(string)
	orderID := c.Param("orderID")

	removeRequest := orders.RemoveOrderRequest{
		OrderID: orderID,
		UserID:  userID,
	}

	r, _ := controller.Orders.RemoveOrder(context.Background(), &removeRequest)
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

	response := &ResponseOrderSuccess{
		Status: "success",
	}

	return c.JSON(http.StatusOK, response)
}
