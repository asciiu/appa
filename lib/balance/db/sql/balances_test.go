package sql_test

import (
	"testing"

	balanceRepo "github.com/asciiu/appa/lib/balance/db/sql"
	balance "github.com/asciiu/appa/lib/balance/models"
	db "github.com/asciiu/appa/lib/db/sql"
	userRepo "github.com/asciiu/appa/lib/user/db/sql"
	user "github.com/asciiu/appa/lib/user/models"
	"github.com/asciiu/appa/lib/util"
	"github.com/stretchr/testify/assert"
)

var testDB string = "postgres://postgres@localhost/appa_test?&sslmode=disable"

func TestInsertBalance(t *testing.T) {
	db, err := db.NewDB(testDB)
	util.CheckErr(err)
	defer db.Close()

	newUser := user.NewUser("flowtester", "test@email", "Yo yo yo!!")
	err = userRepo.InsertUser(db, newUser)
	assert.Nil(t, err, "insert new user failed")

	amount := int64(1000000000000)
	locked := int64(0)
	symbol := "BTC"
	name := "Bitcoin"
	address := "1234"
	precision := 0.000000000001
	newBalance := balance.NewBalance(newUser.ID, symbol, name, address, amount, locked)
	err = balanceRepo.InsertBalance(db, newBalance)

	assert.Nil(t, err, "insert balance failed")

	foundBalance, err := balanceRepo.FindUserBalanceBySymbol(db, newUser.ID, "BTC")

	assert.Equal(t, symbol, foundBalance.Symbol, "symbol should be BTC")
	assert.Equal(t, name, foundBalance.Name, "name should be Bitcoin")
	assert.Equal(t, address, foundBalance.Address, "address did not match")
	assert.Equal(t, balance.Int64(amount), foundBalance.Amount, "amount did not match")
	assert.Equal(t, balance.Int64(locked), foundBalance.Locked, "locked did not match")
	assert.Equal(t, precision, foundBalance.Precision, "precision does not match")

	userRepo.DeleteUserHard(db, newUser.ID)
}

// Should fail to insert balance when currency name does not exist
func TestFailInsertBalance(t *testing.T) {
	db, err := db.NewDB(testDB)
	util.CheckErr(err)
	defer db.Close()

	newUser := user.NewUser("flowtester", "test@email", "Yo yo yo!!")
	err = userRepo.InsertUser(db, newUser)
	assert.Nil(t, err, "insert new user failed")

	amount := int64(1000000000000)
	locked := int64(0)
	symbol := "ABC"
	name := "LALA"
	address := "1234"
	balance := balance.NewBalance(newUser.ID, symbol, name, address, amount, locked)
	err = balanceRepo.InsertBalance(db, balance)

	assert.NotNil(t, err, "insert balance failed")

	userRepo.DeleteUserHard(db, newUser.ID)
}

func TestFindBalances(t *testing.T) {
	db, err := db.NewDB(testDB)
	util.CheckErr(err)
	defer db.Close()

	newUser := user.NewUser("flowtester", "test@email", "Yo yo yo!!")
	err = userRepo.InsertUser(db, newUser)
	assert.Nil(t, err, "insert new user failed")

	btc := balance.NewBalance(newUser.ID, "BTC", "Bitcoin", "", 1, 0)
	ltc := balance.NewBalance(newUser.ID, "LTC", "Litcoin", "", 2, 0)
	err = balanceRepo.InsertBalance(db, btc)
	assert.Nil(t, err, "insert balance failed")

	err = balanceRepo.InsertBalance(db, ltc)
	assert.Nil(t, err, "insert balance failed")

	balances, err := balanceRepo.FindUserBalances(db, newUser.ID)
	assert.Nil(t, err, "find balances failed")
	assert.Equal(t, 2, len(balances), "must be 2 balances")
	userRepo.DeleteUserHard(db, newUser.ID)
}
