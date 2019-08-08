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

	pagedTrades, err := repoTrade.FindUserTrades(db, user.ID, 0, 100)
	assert.Nil(t, err, "find trades should be nil")
	assert.Equal(t, uint32(1), pagedTrades.Total, "should be a single trade")
	assert.Equal(t, uint32(0), pagedTrades.Page, "page incorrect")
	assert.Equal(t, uint32(100), pagedTrades.PageSize, "page size incorrect")

	repoUser.DeleteUserHard(db, user.ID)
}
