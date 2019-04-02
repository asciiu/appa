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
	precision := 12
	balance := models.NewBalance(user.ID, symbol, name, address, amount, locked, precision)
	err = sql.InsertBalance(db, balance)

	assert.Nil(t, err, "insert balance failed")

	foundBalance, err := sql.FindUserBalanceBySymbol(db, user.ID, "BTC")

	assert.Equal(t, symbol, foundBalance.Symbol, "symbol should be BTC")
	assert.Equal(t, name, foundBalance.Name, "name should be Bitcoin")
	assert.Equal(t, address, foundBalance.Address, "address did not match")
	assert.Equal(t, amount, foundBalance.Amount, "amount did not match")
	assert.Equal(t, locked, foundBalance.Locked, "locked did not match")
	assert.Equal(t, precision, foundBalance.Precision, "precision did not match")

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
	precision := 12
	balance := models.NewBalance(user.ID, symbol, name, address, amount, locked, precision)
	err = sql.InsertBalance(db, balance)

	assert.NotNil(t, err, "insert balance failed")

	sql.DeleteUserHard(db, user.ID)
}
