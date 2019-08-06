package sql_test

import (
	"log"
	"testing"

	repoUser "github.com/asciiu/appa/api-graphql/db/sql"
	user "github.com/asciiu/appa/api-graphql/models"
	"github.com/asciiu/appa/common/db"
	"github.com/asciiu/appa/trade-engine/constants"
	repoTrade "github.com/asciiu/appa/trade-engine/db/sql"
	"github.com/asciiu/appa/trade-engine/models"

	"github.com/stretchr/testify/assert"
)

func TestInsertTrade(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()

	user := user.NewUser("testy", "test@email", "pass")
	err = repoUser.InsertUser(db, user)
	assert.Nil(t, err, "insert new user should be nil")

	sellOrder := models.NewOrder(user.ID, "ada-btc", constants.Sell, 1000, 500)
	buyOrder := models.NewOrder(user.ID, "ada-btc", constants.Buy, 1000, 500)

	err = repoTrade.InsertOrder(db, sellOrder)
	assert.Equal(t, nil, err, "err should be nil")

	err = repoTrade.InsertOrder(db, buyOrder)
	assert.Equal(t, nil, err, "err should be nil")

	newTrade := models.NewTrade(buyOrder.ID, sellOrder.ID, constants.Buy, 1000, 500)
	err = repoTrade.InsertTrade(db, newTrade)
	assert.Equal(t, nil, err, "err should be nil")

	repoUser.DeleteUserHard(db, user.ID)
}
