package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/asciiu/appa/common/db"
	constOrder "github.com/asciiu/appa/order-service/constants"
	protoOrder "github.com/asciiu/appa/order-service/proto/order"
	repoUser "github.com/asciiu/appa/user-service/db/sql"
	user "github.com/asciiu/appa/user-service/models"
	"github.com/stretchr/testify/assert"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func setupService() (*OrderService, *user.User) {
	dbUrl := "postgres://postgres@localhost:5432/appa_test?&sslmode=disable"
	db, _ := db.NewDB(dbUrl)

	orderService := OrderService{DB: db}

	user := user.NewUser("first", "last", "test@email", "hash")
	_, err := repoUser.InsertUser(db, user)
	checkErr(err)

	return &orderService, user
}

func TestAddOrder(t *testing.T) {
	service, user := setupService()

	defer service.DB.Close()

	req := protoOrder.NewOrderRequest{
		UserID:     user.ID,
		MarketName: "test-btc",
		Side:       constOrder.Sell,
		Size:       1.0,
		Type:       constOrder.LimitOrder,
	}

	res := protoOrder.OrderResponse{}
	service.AddOrder(context.Background(), &req, &res)
	assert.Equal(t, "success", res.Status, "expected success got: "+res.Message)

	repoUser.DeleteUserHard(service.DB, user.ID)
}

func TestFindOrder(t *testing.T) {
	service, user := setupService()

	defer service.DB.Close()

	req := protoOrder.NewOrderRequest{
		UserID:     user.ID,
		MarketName: "test-btc",
		Side:       constOrder.Sell,
		Size:       1.0,
		Type:       constOrder.LimitOrder,
	}

	res := protoOrder.OrderResponse{}
	service.AddOrder(context.Background(), &req, &res)
	assert.Equal(t, "success", res.Status, "expected success got: "+res.Message)
	order := res.Data.Order

	req2 := protoOrder.OrderRequest{
		OrderID: order.OrderID,
		UserID:  user.ID,
	}

	res2 := protoOrder.OrderResponse{}
	service.FindOrder(context.Background(), &req2, &res2)
	assert.Equal(t, "success", res2.Status, "expected success got: "+res2.Message)
	assert.Equal(t, order.OrderID, res2.Data.Order.OrderID, "order ID in find does not match")

	repoUser.DeleteUserHard(service.DB, user.ID)
}

func TestFindOrderWrongUserID(t *testing.T) {
	service, user := setupService()

	defer service.DB.Close()

	req := protoOrder.NewOrderRequest{
		UserID:     user.ID,
		MarketName: "test-btc",
		Side:       constOrder.Sell,
		Size:       1.0,
		Type:       constOrder.LimitOrder,
	}

	res := protoOrder.OrderResponse{}
	service.AddOrder(context.Background(), &req, &res)
	assert.Equal(t, "success", res.Status, "expected success got: "+res.Message)
	order := res.Data.Order

	req2 := protoOrder.OrderRequest{
		OrderID: order.OrderID,
		UserID:  "a1c0e0dd-0c73-4b5e-ac5b-a2ac8378427d",
	}

	res2 := protoOrder.OrderResponse{}
	service.FindOrder(context.Background(), &req2, &res2)
	assert.Equal(t, "nonentity", res2.Status, "expected nonentity got: "+res2.Message)

	repoUser.DeleteUserHard(service.DB, user.ID)
}

func TestCancelOrder(t *testing.T) {
	service, user := setupService()

	defer service.DB.Close()

	req := protoOrder.NewOrderRequest{
		UserID:     user.ID,
		MarketName: "test-btc",
		Side:       constOrder.Sell,
		Size:       1.0,
		Type:       constOrder.LimitOrder,
	}

	res := protoOrder.OrderResponse{}
	service.AddOrder(context.Background(), &req, &res)
	assert.Equal(t, "success", res.Status, "expected success got: "+res.Message)
	order := res.Data.Order

	req2 := protoOrder.OrderRequest{
		OrderID: order.OrderID,
		UserID:  user.ID,
	}

	res2 := protoOrder.StatusResponse{}
	service.CancelOrder(context.Background(), &req2, &res2)
	assert.Equal(t, "success", res2.Status, "expected success got: "+res2.Message)

	req3 := protoOrder.OrderRequest{
		OrderID: order.OrderID,
		UserID:  user.ID,
	}

	res3 := protoOrder.OrderResponse{}
	service.FindOrder(context.Background(), &req3, &res3)
	assert.Equal(t, "nonentity", res3.Status, "expected nonentity got: "+res3.Message)

	repoUser.DeleteUserHard(service.DB, user.ID)
}

func TestCancelOrderWrongUserID(t *testing.T) {
	service, user := setupService()

	defer service.DB.Close()

	req := protoOrder.NewOrderRequest{
		UserID:     user.ID,
		MarketName: "test-btc",
		Side:       constOrder.Sell,
		Size:       1.0,
		Type:       constOrder.LimitOrder,
	}

	res := protoOrder.OrderResponse{}
	service.AddOrder(context.Background(), &req, &res)
	assert.Equal(t, "success", res.Status, "expected success got: "+res.Message)
	order := res.Data.Order

	req2 := protoOrder.OrderRequest{
		OrderID: order.OrderID,
		UserID:  "a1c0e0dd-0c73-4b5e-ac5b-a2ac8378427d",
	}

	res2 := protoOrder.StatusResponse{}
	service.CancelOrder(context.Background(), &req2, &res2)
	assert.Equal(t, "success", res2.Status, "expected success got: "+res2.Message)

	req3 := protoOrder.OrderRequest{
		OrderID: order.OrderID,
		UserID:  user.ID,
	}

	res3 := protoOrder.OrderResponse{}
	service.FindOrder(context.Background(), &req3, &res3)
	assert.Equal(t, "success", res3.Status, "expected success got: "+res3.Message)
	assert.Equal(t, order.OrderID, res3.Data.Order.OrderID, "order IDs do not match")

	repoUser.DeleteUserHard(service.DB, user.ID)
}

func TestFindUserOrders(t *testing.T) {
	service, user := setupService()

	defer service.DB.Close()

	for i := 0; i < 10; i++ {
		req := protoOrder.NewOrderRequest{
			UserID:     user.ID,
			MarketName: "test-btc",
			Side:       constOrder.Sell,
			Size:       1.0,
			Type:       constOrder.LimitOrder,
		}

		res := protoOrder.OrderResponse{}
		service.AddOrder(context.Background(), &req, &res)
		assert.Equal(t, "success", res.Status, "expected success got: "+res.Message)
	}

	req2 := protoOrder.UserOrdersRequest{
		UserID:   user.ID,
		Page:     0,
		PageSize: 20,
		Status:   "",
	}

	res2 := protoOrder.OrdersPageResponse{}
	service.FindUserOrders(context.Background(), &req2, &res2)
	assert.Equal(t, "success", res2.Status, "expected success got: "+res2.Message)
	assert.Equal(t, 10, len(res2.Data.Orders), fmt.Sprintf("must be 10 orders got %d", len(res2.Data.Orders)))

	repoUser.DeleteUserHard(service.DB, user.ID)
}
