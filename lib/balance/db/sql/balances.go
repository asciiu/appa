package sql

import (
	"database/sql"

	balance "github.com/asciiu/appa/lib/balance/models"
)

func DeleteBalance(db *sql.DB, balanceID string) error {
	_, err := db.Exec("DELETE FROM balances WHERE id = $1", balanceID)
	return err
}

func FindCurrency(db *sql.DB, symbol string) (*balance.Currency, error) {
	var c balance.Currency
	err := db.QueryRow(`SELECT 
	symbol,
	name,
    precision
	FROM currencies
	WHERE symbol = $1`, symbol).
		Scan(&c.Symbol,
			&c.Name,
			&c.Precision)

	if err != nil {
		return nil, err
	}
	return &c, nil
}

func FindUserBalanceBySymbol(db *sql.DB, userID, symbol string) (*balance.Balance, error) {
	var b balance.Balance
	err := db.QueryRow(`SELECT 
	b.id, 
	b.user_id, 
	b.symbol,
	c.name,
	b.amount,
	b.locked,
	c.precision,
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

func FindUserBalances(db *sql.DB, userID string) ([]*balance.Balance, error) {

	rows, err := db.Query(`SELECT 
	    b.id, 
	    b.user_id, 
	    b.symbol,
	    c.name,
	    b.amount,
		b.locked,
		c.precision,
	    b.address 
	    FROM balances b
	    JOIN currencies c ON b.symbol = c.symbol
	    WHERE b.user_id = $1`, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	balances := make([]*balance.Balance, 0)
	for rows.Next() {
		b := new(balance.Balance)
		err := rows.Scan(&b.ID,
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
		balances = append(balances, b)
	}

	return balances, nil
}

func InsertBalance(db *sql.DB, bal *balance.Balance) error {
	sqlStatement := `insert into balances (
	    id, 
	    user_id, 
	    symbol,
	    amount,
	    locked,
	    address) values ($1, $2, $3, $4, $5, $6)`

	_, err := db.Exec(sqlStatement,
		bal.ID,
		bal.UserID,
		bal.Symbol,
		bal.Amount,
		bal.Locked,
		bal.Address)

	return err
}
