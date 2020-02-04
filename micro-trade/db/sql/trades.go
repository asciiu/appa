package sql

import (
	"database/sql"

	"github.com/asciiu/appa/trade-engine/models"
)

func InsertTrade(db *sql.DB, trade *models.Trade) error {
	sqlStatement := `insert into trades (
		id,
		taker_order_id, 
		maker_order_id, 
		amount,
		price,
		side, 
		created_on, 
		updated_on) values ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := db.Exec(sqlStatement,
		trade.ID,
		trade.TakerOrderID,
		trade.MakerOrderID,
		trade.Amount,
		trade.Price,
		trade.Side,
		trade.CreatedOn,
		trade.UpdatedOn,
	)

	return err
}

func FindUserTrades(db *sql.DB, userID string, page, pageSize uint32) (*models.TradesPage, error) {
	pagedTrades := &models.TradesPage{
		Page:     page,
		PageSize: pageSize,
		Trades:   make([]*models.Trade, 0),
	}

	queryCount := `SELECT count(distinct t.id) FROM orders o 
				   JOIN trades t ON t.taker_order_id = o.id OR t.maker_order_id = o.id
	               WHERE o.user_id = $1`
	err := db.QueryRow(queryCount, userID).Scan(&pagedTrades.Total)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`SELECT 
	    t.id,
		t.taker_order_id, 
		t.maker_order_id,
	    t.side, 
		t.amount, 
		t.price,
		t.created_on,
		t.updated_on
		FROM orders o 
		JOIN trades t ON t.taker_order_id = o.id OR t.maker_order_id = o.id
		WHERE user_id = $1 OFFSET $2 LIMIT $3`, userID, page, pageSize)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		t := new(models.Trade)
		err := rows.Scan(
			&t.ID,
			&t.TakerOrderID,
			&t.MakerOrderID,
			&t.Side,
			&t.Amount,
			&t.Price,
			&t.CreatedOn,
			&t.UpdatedOn,
		)

		if err != nil {
			return nil, err
		}
		pagedTrades.Trades = append(pagedTrades.Trades, t)
	}

	return pagedTrades, nil
}
