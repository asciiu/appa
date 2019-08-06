package main

import (
	"log"

	repo "github.com/asciiu/appa/api-graphql/db/sql"
	user "github.com/asciiu/appa/api-graphql/models"
	"github.com/asciiu/appa/common/db"
	"github.com/asciiu/appa/trade-engine/models"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func setupService() (*TradeEngine, *user.User) {
	dbUrl := "postgres://postgres@localhost:5432/appa_test?&sslmode=disable"
	db, _ := db.NewDB(dbUrl)

	engine := TradeEngine{
		DB:         db,
		OrderBooks: make(map[string]*models.OrderBook),
	}

	user := user.NewUser("testy", "test@email", "hash")
	err := repo.InsertUser(db, user)
	checkErr(err)

	return &engine, user
}

//func TestAddOrder(t *testing.T) {
//	service, user := setupService()
//
//	defer service.DB.Close()
//
//	req := trade.NewOrderRequest{
//		UserID:     user.ID,
//		MarketName: "test-btc",
//		Side:       constants.Sell,
//		Size:       1.0,
//		Type:       constants.LimitOrder,
//	}
//
//	res := trade.OrderResponse{}
//	service.AddOrder(context.Background(), &req, &res)
//	assert.Equal(t, "success", res.Status, "expected success got: "+res.Message)
//
//	sql.DeleteUserHard(service.DB, user.ID)
//}
//
//func TestFindOrder(t *testing.T) {
//	service, user := setupService()
//
//	defer service.DB.Close()
//
//	req := trade.NewOrderRequest{
//		UserID:     user.ID,
//		MarketName: "test-btc",
//		Side:       constants.Sell,
//		Size:       1.0,
//		Type:       constants.LimitOrder,
//	}
//
//	res := trade.OrderResponse{}
//	service.AddOrder(context.Background(), &req, &res)
//	assert.Equal(t, "success", res.Status, "expected success got: "+res.Message)
//	order := res.Data.Order
//
//	req2 := trade.OrderRequest{
//		OrderID: order.OrderID,
//		UserID:  user.ID,
//	}
//
//	res2 := trade.OrderResponse{}
//	service.FindOrder(context.Background(), &req2, &res2)
//	assert.Equal(t, "success", res2.Status, "expected success got: "+res2.Message)
//	assert.Equal(t, order.OrderID, res2.Data.Order.OrderID, "order ID in find does not match")
//
//	sql.DeleteUserHard(service.DB, user.ID)
//}
//
//func TestFindOrderWrongUserID(t *testing.T) {
//	service, user := setupService()
//
//	defer service.DB.Close()
//
//	req := trade.NewOrderRequest{
//		UserID:     user.ID,
//		MarketName: "test-btc",
//		Side:       constants.Sell,
//		Size:       1.0,
//		Type:       constants.LimitOrder,
//	}
//
//	res := trade.OrderResponse{}
//	service.AddOrder(context.Background(), &req, &res)
//	assert.Equal(t, "success", res.Status, "expected success got: "+res.Message)
//	order := res.Data.Order
//
//	req2 := trade.OrderRequest{
//		OrderID: order.OrderID,
//		UserID:  "a1c0e0dd-0c73-4b5e-ac5b-a2ac8378427d",
//	}
//
//	res2 := trade.OrderResponse{}
//	service.FindOrder(context.Background(), &req2, &res2)
//	assert.Equal(t, "nonentity", res2.Status, "expected nonentity got: "+res2.Message)
//
//	sql.DeleteUserHard(service.DB, user.ID)
//}
//
//func TestCancelOrder(t *testing.T) {
//	service, user := setupService()
//
//	defer service.DB.Close()
//
//	req := trade.NewOrderRequest{
//		UserID:     user.ID,
//		MarketName: "test-btc",
//		Side:       constants.Sell,
//		Size:       1.0,
//		Type:       constants.LimitOrder,
//	}
//
//	res := trade.OrderResponse{}
//	service.AddOrder(context.Background(), &req, &res)
//	assert.Equal(t, "success", res.Status, "expected success got: "+res.Message)
//	order := res.Data.Order
//
//	req2 := trade.OrderRequest{
//		OrderID: order.OrderID,
//		UserID:  user.ID,
//	}
//
//	res2 := trade.StatusResponse{}
//	service.CancelOrder(context.Background(), &req2, &res2)
//	assert.Equal(t, "success", res2.Status, "expected success got: "+res2.Message)
//
//	req3 := trade.OrderRequest{
//		OrderID: order.OrderID,
//		UserID:  user.ID,
//	}
//
//	res3 := trade.OrderResponse{}
//	service.FindOrder(context.Background(), &req3, &res3)
//	assert.Equal(t, "nonentity", res3.Status, "expected nonentity got: "+res3.Message)
//
//	sql.DeleteUserHard(service.DB, user.ID)
//}
//
//func TestCancelOrderWrongUserID(t *testing.T) {
//	service, user := setupService()
//
//	defer service.DB.Close()
//
//	req := trade.NewOrderRequest{
//		UserID:     user.ID,
//		MarketName: "test-btc",
//		Side:       constants.Sell,
//		Size:       1.0,
//		Type:       constants.LimitOrder,
//	}
//
//	res := trade.OrderResponse{}
//	service.AddOrder(context.Background(), &req, &res)
//	assert.Equal(t, "success", res.Status, "expected success got: "+res.Message)
//	order := res.Data.Order
//
//	req2 := trade.OrderRequest{
//		OrderID: order.OrderID,
//		UserID:  "a1c0e0dd-0c73-4b5e-ac5b-a2ac8378427d",
//	}
//
//	res2 := trade.StatusResponse{}
//	service.CancelOrder(context.Background(), &req2, &res2)
//	assert.Equal(t, "nonentity", res2.Status, "expected nonentity got: "+res2.Message)
//
//	req3 := trade.OrderRequest{
//		OrderID: order.OrderID,
//		UserID:  user.ID,
//	}
//
//	res3 := trade.OrderResponse{}
//	service.FindOrder(context.Background(), &req3, &res3)
//	assert.Equal(t, "success", res3.Status, "expected success got: "+res3.Message)
//	assert.Equal(t, order.OrderID, res3.Data.Order.OrderID, "order IDs do not match")
//
//	sql.DeleteUserHard(service.DB, user.ID)
//}
//
//func TestFindUserOrders(t *testing.T) {
//	service, user := setupService()
//
//	defer service.DB.Close()
//
//	for i := 0; i < 10; i++ {
//		req := trade.NewOrderRequest{
//			UserID:     user.ID,
//			MarketName: "test-btc",
//			Side:       constants.Sell,
//			Size:       1.0,
//			Type:       constants.LimitOrder,
//		}
//
//		res := trade.OrderResponse{}
//		service.AddOrder(context.Background(), &req, &res)
//		assert.Equal(t, "success", res.Status, "expected success got: "+res.Message)
//	}
//
//	req2 := trade.UserOrdersRequest{
//		UserID:   user.ID,
//		Page:     0,
//		PageSize: 20,
//		Status:   "",
//	}
//
//	res2 := trade.OrdersPageResponse{}
//	service.FindUserOrders(context.Background(), &req2, &res2)
//	assert.Equal(t, "success", res2.Status, "expected success got: "+res2.Message)
//	assert.Equal(t, 10, len(res2.Data.Orders), fmt.Sprintf("must be 10 orders got %d", len(res2.Data.Orders)))
//
//	sql.DeleteUserHard(service.DB, user.ID)
//}
