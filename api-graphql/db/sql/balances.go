package sql

import (
	"database/sql"

	"github.com/asciiu/appa/api-graphql/models"
)

func DeleteBalance(db *sql.DB, balanceID string) error {
	_, err := db.Exec("DELETE FROM balances WHERE id = $1", balanceID)
	return err
}

func FindCurrency(db *sql.DB, symbol string) (*models.Currency, error) {
	var c models.Currency
	err := db.QueryRow(`SELECT 
	symbol,
	name
	FROM currencies
	WHERE symbol = $1`, symbol).
		Scan(&c.Symbol,
			&c.Name)

	if err != nil {
		return nil, err
	}
	return &c, nil
}

func FindUserBalanceBySymbol(db *sql.DB, userID, symbol string) (*models.Balance, error) {
	var b models.Balance
	err := db.QueryRow(`SELECT 
	b.id, 
	b.user_id, 
	b.symbol,
	c.name,
	b.amount,
	b.locked,
	b.precision,
	b.address 
	FROM balances b
	JOIN currencies c ON b.symbol = c.symbol
	WHERE b.user_id = $1 AND b.symbol = $2`, userID, symbol).
		Scan(&b.ID,
			&b.UserID,
			&b.Symbol,
			&b.Name,
			&b.Amount,
			&b.Locked,
			&b.Precision,
			&b.Address)

	if err != nil {
		return nil, err
	}
	return &b, nil
}

func InsertBalance(db *sql.DB, balance *models.Balance) error {
	sqlStatement := `insert into balances (
	    id, 
	    user_id, 
	    symbol,
	    amount,
	    locked,
	    precision,
	    address) values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.Exec(sqlStatement,
		balance.ID,
		balance.UserID,
		balance.Symbol,
		balance.Amount,
		balance.Locked,
		balance.Precision,
		balance.Address)

	return err
}
