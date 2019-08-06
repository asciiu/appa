package sql_test

import (
	"log"
	"testing"

	repoUser "github.com/asciiu/appa/api-graphql/db/sql"
	user "github.com/asciiu/appa/api-graphql/models"
	"github.com/asciiu/appa/common/db"
	"github.com/asciiu/appa/order-service/constants"
	repoOrder "github.com/asciiu/appa/order-service/db/sql"
	"github.com/asciiu/appa/order-service/models"

	"github.com/stretchr/testify/assert"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func TestInsertOrder(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	user := user.NewUser("testy", "test@email", "pass")
	err = repoUser.InsertUser(db, user)
	assert.Nil(t, err, "insert new user should be nil")

	newOrder := models.NewOrder(user.ID, "ada-btc", constants.Sell, 1000, 500)
	err = repoOrder.InsertOrder(db, newOrder)
	assert.Equal(t, nil, err, "err should be nil")

	repoUser.DeleteUserHard(db, user.ID)
}

func TestFindOrder(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	user := user.NewUser("testy2", "test@email", "pass")
	err = repoUser.InsertUser(db, user)
	assert.Nil(t, err, "insert new user should be nil")

	newOrder := models.NewOrder(user.ID, "test-btc", constants.Buy, 2, 100)
	err = repoOrder.InsertOrder(db, newOrder)
	assert.Equal(t, nil, err, "err should be nil")

	findOrder, err := repoOrder.FindOrderByID(db, newOrder.ID)
	assert.Equal(t, nil, err, "err should be nil")

	assert.Equal(t, newOrder.ID, findOrder.ID, "order ids not equal")
	assert.Equal(t, newOrder.UserID, findOrder.UserID, "user ids not equal")
	assert.Equal(t, newOrder.MarketName, findOrder.MarketName, "market names not equal")
	assert.Equal(t, newOrder.Side, findOrder.Side, "sides not equal")
	assert.Equal(t, newOrder.Amount, findOrder.Amount, "amount not equal")
	assert.Equal(t, newOrder.Filled, findOrder.Filled, "fills not equal")
	assert.Equal(t, newOrder.Type, findOrder.Type, "type not equal")
	assert.Equal(t, newOrder.Status, findOrder.Status, "status not equal")

	repoUser.DeleteUserHard(db, user.ID)
}

func TestFindUserOrders(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	user := user.NewUser("testy2", "test@email", "pass")
	err = repoUser.InsertUser(db, user)
	assert.Nil(t, err, "insert new user should be nil")

	newOrder1 := models.NewOrder(user.ID, "test-btc", constants.Buy, 2, 100)
	err = repoOrder.InsertOrder(db, newOrder1)
	assert.Equal(t, nil, err, "err should be nil")

	newOrder2 := models.NewOrder(user.ID, "test-btc", constants.Buy, 2, 100)
	err = repoOrder.InsertOrder(db, newOrder2)
	assert.Equal(t, nil, err, "err should be nil")

	page, err := repoOrder.FindUserOrders(db, user.ID, constants.Pending, 0, 10)
	assert.Equal(t, nil, err, "err should be nil")

	assert.Equal(t, uint32(0), page.Page, "page not expected")
	assert.Equal(t, uint32(10), page.PageSize, "page size not expected")
	assert.Equal(t, uint32(2), page.Total, "should be total of 2")
	assert.Equal(t, 2, len(page.Orders), "should be 2 orders")

	repoUser.DeleteUserHard(db, user.ID)
}

func TestUpdateOrderStatus(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	user := user.NewUser("testy2", "test@email", "pass")
	err = repoUser.InsertUser(db, user)
	assert.Nil(t, err, "insert new user should be nil")

	newOrder := models.NewOrder(user.ID, "test-btc", constants.Buy, 2, 100)
	err = repoOrder.InsertOrder(db, newOrder)
	assert.Equal(t, nil, err, "err should be nil")

	repoOrder.UpdateOrderStatus(db, newOrder.ID, constants.Cancelled)

	findOrder, err := repoOrder.FindOrderByID(db, newOrder.ID)
	assert.Equal(t, nil, err, "err should be nil")

	assert.Equal(t, newOrder.ID, findOrder.ID, "order ids not equal")
	assert.Equal(t, constants.Cancelled, findOrder.Status, "status should be cancelled")

	repoUser.DeleteUserHard(db, user.ID)
}

func TestUpdateOrderAmounts(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	user := user.NewUser("testy2", "test@email", "pass")
	err = repoUser.InsertUser(db, user)
	assert.Nil(t, err, "insert new user should be nil")

	newOrder := models.NewOrder(user.ID, "test-btc", constants.Buy, 2, 100)
	err = repoOrder.InsertOrder(db, newOrder)
	assert.Equal(t, nil, err, "err should be nil")

	repoOrder.UpdateOrderAmounts(db, newOrder.ID, 1, 1)

	findOrder, err := repoOrder.FindOrderByID(db, newOrder.ID)
	assert.Equal(t, nil, err, "err should be nil")

	assert.Equal(t, newOrder.ID, findOrder.ID, "order ids not equal")
	assert.Equal(t, uint64(1), findOrder.Amount, "should be 1 for amount")
	assert.Equal(t, uint64(1), findOrder.Filled, "should be 1 for filled")

	repoUser.DeleteUserHard(db, user.ID)
}
