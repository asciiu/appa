package sql_test

import (
	"log"
	"testing"
	"time"

	"github.com/asciiu/appa/common/db"
	"github.com/asciiu/appa/order-service/constants"
	repoOrder "github.com/asciiu/appa/order-service/db/sql"
	protoOrder "github.com/asciiu/appa/order-service/proto/order"
	repoUser "github.com/asciiu/appa/user-service/db/sql"
	"github.com/asciiu/appa/user-service/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
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

	user := models.NewUser("first", "last", "test@email", "hash")
	_, error := repoUser.InsertUser(db, user)
	if error != nil {
		t.Errorf("%s", error)
	}

	now := string(pq.FormatTimestamp(time.Now().UTC()))
	newOrder := protoOrder.Order{
		OrderID:    uuid.New().String(),
		UserID:     user.ID,
		MarketName: "ada-btc",
		Side:       constants.Sell,
		Size:       0.00001,
		Type:       constants.LimitOrder,
		Status:     constants.Pending,
		CreatedOn:  now,
		UpdatedOn:  now,
	}

	savedOrder, _ := repoOrder.InsertOrder(db, &newOrder)
	assert.Equal(t, newOrder.OrderID, savedOrder.OrderID, "order ids not equal")
	assert.Equal(t, newOrder.UserID, savedOrder.UserID, "user ids not equal")
	assert.Equal(t, newOrder.MarketName, savedOrder.MarketName, "market names not equal")
	assert.Equal(t, newOrder.Side, savedOrder.Side, "sides not equal")
	assert.Equal(t, newOrder.Size, savedOrder.Size, "size not equal")
	assert.Equal(t, newOrder.Type, savedOrder.Type, "type not equal")
	assert.Equal(t, newOrder.Status, savedOrder.Status, "status not equal")

	repoUser.DeleteUserHard(db, user.ID)
}

func TestFindOrder(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	user := models.NewUser("first", "last", "test@email", "hash")
	_, error := repoUser.InsertUser(db, user)
	if error != nil {
		t.Errorf("%s", error)
	}

	now := string(pq.FormatTimestamp(time.Now().UTC()))
	newOrder := protoOrder.Order{
		OrderID:    uuid.New().String(),
		UserID:     user.ID,
		MarketName: "test-btc",
		Side:       constants.Buy,
		Size:       1.0,
		Type:       constants.LimitOrder,
		Status:     constants.Pending,
		CreatedOn:  now,
		UpdatedOn:  now,
	}

	repoOrder.InsertOrder(db, &newOrder)
	findOrder, _ := repoOrder.FindOrder(db, newOrder.OrderID)

	assert.Equal(t, newOrder.OrderID, findOrder.OrderID, "order ids not equal")
	assert.Equal(t, newOrder.UserID, findOrder.UserID, "user ids not equal")
	assert.Equal(t, newOrder.MarketName, findOrder.MarketName, "market names not equal")
	assert.Equal(t, newOrder.Side, findOrder.Side, "sides not equal")
	assert.Equal(t, newOrder.Size, findOrder.Size, "size not equal")
	assert.Equal(t, newOrder.Type, findOrder.Type, "type not equal")
	assert.Equal(t, newOrder.Status, findOrder.Status, "status not equal")

	repoUser.DeleteUserHard(db, user.ID)
}
