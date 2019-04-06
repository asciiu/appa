package sql_test

import (
	"testing"

	"github.com/asciiu/appa/api-graphql/db/sql"
	"github.com/asciiu/appa/api-graphql/models"
	"github.com/asciiu/appa/common/db"
	"github.com/stretchr/testify/assert"
)

func TestInsertBalance(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	user := models.NewUser("flowtester", "test@email", "Yo yo yo!!")
	err = sql.InsertUser(db, user)
	assert.Nil(t, err, "insert new user failed")

	amount := int64(1000000000000)
	locked := int64(0)
	symbol := "BTC"
	name := "Bitcoin"
	address := "1234"
	precision := 0.000000000001
	balance := models.NewBalance(user.ID, symbol, name, address, amount, locked)
	err = sql.InsertBalance(db, balance)

	assert.Nil(t, err, "insert balance failed")

	foundBalance, err := sql.FindUserBalanceBySymbol(db, user.ID, "BTC")

	assert.Equal(t, symbol, foundBalance.Symbol, "symbol should be BTC")
	assert.Equal(t, name, foundBalance.Name, "name should be Bitcoin")
	assert.Equal(t, address, foundBalance.Address, "address did not match")
	assert.Equal(t, models.Int64(amount), foundBalance.Amount, "amount did not match")
	assert.Equal(t, models.Int64(locked), foundBalance.Locked, "locked did not match")
	assert.Equal(t, precision, foundBalance.Precision, "precision does not match")

	sql.DeleteUserHard(db, user.ID)
}

// Should fail to insert balance when currency name does not exist
func TestFailInsertBalance(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	user := models.NewUser("flowtester", "test@email", "Yo yo yo!!")
	err = sql.InsertUser(db, user)
	assert.Nil(t, err, "insert new user failed")

	amount := int64(1000000000000)
	locked := int64(0)
	symbol := "ABC"
	name := "LALA"
	address := "1234"
	balance := models.NewBalance(user.ID, symbol, name, address, amount, locked)
	err = sql.InsertBalance(db, balance)

	assert.NotNil(t, err, "insert balance failed")

	sql.DeleteUserHard(db, user.ID)
}

func TestFindBalances(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	user := models.NewUser("flowtester", "test@email", "Yo yo yo!!")
	err = sql.InsertUser(db, user)
	assert.Nil(t, err, "insert new user failed")

	btc := models.NewBalance(user.ID, "BTC", "Bitcoin", "", 1, 0)
	ltc := models.NewBalance(user.ID, "LTC", "Litcoin", "", 2, 0)
	err = sql.InsertBalance(db, btc)
	assert.Nil(t, err, "insert balance failed")

	err = sql.InsertBalance(db, ltc)
	assert.Nil(t, err, "insert balance failed")

	balances, err := sql.FindUserBalances(db, user.ID)
	assert.Nil(t, err, "find balances failed")
	assert.Equal(t, 2, len(balances), "must be 2 balances")
	sql.DeleteUserHard(db, user.ID)
}
