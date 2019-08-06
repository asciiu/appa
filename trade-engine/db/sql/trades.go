package sql

import (
	"database/sql"

	"github.com/asciiu/appa/trade-engine/models"
)

func InsertTrade(db *sql.DB, trade *models.Trade) error {
	sqlStatement := `insert into trades (
		taker_order_id, 
		maker_order_id, 
		amount,
		price,
		side, 
		created_on, 
		updated_on) values ($1, $2, $3, $4, $5, $6, $7)`
	_, err := db.Exec(sqlStatement,
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
