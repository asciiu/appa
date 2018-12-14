package sql

import (
	"database/sql"

	"github.com/asciiu/oldiez/order-service/models"
)

func DeleteOrder(db *sql.DB, orderID string) error {
	_, err := db.Exec("DELETE FROM orders WHERE id = $1", orderID)
	return err
}

func FindOrder(db *sql.DB, orderID string) (*models.Order, error) {
	var o models.Order
	err := db.QueryRow(`SELECT 
	    id, 
	    user_id, 
	    market_name, 
	    side, 
	    size, 
	    created_on 
	    FROM orders WHERE id = $1`, orderID).Scan(
		&o.OrderID,
		&o.UserID,
		&o.MarketName,
		&o.Side,
		&o.Size,
		&o.CreatedOn,
	)

	if err != nil {
		return nil, err
	}
	return &o, nil
}

func InsertOrder(db *sql.DB, order *models.Order) (*models.Order, error) {
	sqlStatement := `insert into orders (
		id, 
		user_id, 
		market_name, 
		side, 
		size, 
		created_on) values ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(sqlStatement,
		order.OrderID,
		order.UserID,
		order.MarketName,
		order.Side,
		order.Size,
		order.CreatedOn,
	)

	if err != nil {
		return nil, err
	}
	return order, nil
}
