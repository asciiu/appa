package sql_test

import (
	"log"
	"testing"
	"time"

	"github.com/asciiu/oldiez/common/db"
	"github.com/asciiu/oldiez/order-service/constants"
	repoOrder "github.com/asciiu/oldiez/order-service/db/sql"
	protoOrder "github.com/asciiu/oldiez/order-service/proto/order"
	repoUser "github.com/asciiu/oldiez/user-service/db/sql"
	"github.com/asciiu/oldiez/user-service/models"
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
	db, err := db.NewDB("postgres://postgres@localhost/oldiez_test?&sslmode=disable")
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
		CreatedOn:  now,
		UpdatedOn:  now,
	}

	savedOrder, _ := repoOrder.InsertOrder(db, &newOrder)
	assert.Equal(t, newOrder.OrderID, savedOrder.OrderID, "order ids not equal")
	assert.Equal(t, newOrder.UserID, savedOrder.UserID, "user ids not equal")
	assert.Equal(t, newOrder.MarketName, savedOrder.MarketName, "market names not equal")
	assert.Equal(t, newOrder.Side, savedOrder.Side, "sides not equal")
	assert.Equal(t, newOrder.Size, savedOrder.Size, "size not equal")

	repoUser.DeleteUserHard(db, user.ID)
}
