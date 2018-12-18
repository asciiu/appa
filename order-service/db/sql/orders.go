package sql

import (
	"database/sql"

	protoOrder "github.com/asciiu/oldiez/order-service/proto/order"
)

func DeleteOrder(db *sql.DB, orderID, userID string) error {
	_, err := db.Exec("DELETE FROM orders WHERE id = $1 and user_id = $2", orderID, userID)
	return err
}

func FindOrder(db *sql.DB, orderID, userID string) (*protoOrder.Order, error) {
	var o protoOrder.Order
	err := db.QueryRow(`SELECT 
	    id, 
	    user_id, 
	    market_name, 
	    side, 
		size, 
		type,
		status,
		created_on,
		updated_on
	    FROM orders WHERE id = $1 and user_id = $2`, orderID, userID).Scan(
		&o.OrderID,
		&o.UserID,
		&o.MarketName,
		&o.Side,
		&o.Size,
		&o.Type,
		&o.Status,
		&o.CreatedOn,
		&o.UpdatedOn,
	)

	if err != nil {
		return nil, err
	}
	return &o, nil
}

func FindUserOrders(db *sql.DB, userID string) ([]*protoOrder.Order, error) {
	rows, err := db.Query(`SELECT 
	    id, 
	    user_id, 
	    market_name, 
	    side, 
		size, 
		type,
		status,
		created_on,
		updated_on
		FROM orders WHERE user_id = $1`, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orders := make([]*protoOrder.Order, 0)
	for rows.Next() {
		o := new(protoOrder.Order)
		err := rows.Scan(
			&o.OrderID,
			&o.UserID,
			&o.MarketName,
			&o.Side,
			&o.Size,
			&o.Type,
			&o.Status,
			&o.CreatedOn,
			&o.UpdatedOn,
		)

		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	return orders, nil
}

func InsertOrder(db *sql.DB, order *protoOrder.Order) error {
	sqlStatement := `insert into orders (
		id, 
		user_id, 
		market_name, 
		side, 
		size, 
		type,
		status,
		created_on, 
		updated_on) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := db.Exec(sqlStatement,
		order.OrderID,
		order.UserID,
		order.MarketName,
		order.Side,
		order.Size,
		order.Type,
		order.Status,
		order.CreatedOn,
		order.UpdatedOn,
	)

	return err
}
