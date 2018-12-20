package sql

import (
	"database/sql"

	protoOrder "github.com/asciiu/appa/order-service/proto/order"
)

func DeleteOrder(db *sql.DB, orderID, userID string) error {
	_, err := db.Exec("DELETE FROM orders WHERE id = $1 and user_id = $2", orderID, userID)
	return err
}

func FindOrder(db *sql.DB, orderID, userID string) (*protoOrder.Order, error) {
	var o protoOrder.Order
	var price sql.NullFloat64
	err := db.QueryRow(`SELECT 
	    id, 
	    user_id, 
	    market_name, 
	    side, 
		size, 
		price,
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
		&price,
		&o.Type,
		&o.Status,
		&o.CreatedOn,
		&o.UpdatedOn,
	)

	if err != nil {
		return nil, err
	}
	if price.Valid {
		o.Price = price.Float64
	}
	return &o, nil
}

func FindUserOrders(db *sql.DB, userID, status string, page, pageSize uint32) (*protoOrder.OrdersPage, error) {
	var total uint32
	queryCount := `SELECT count(*) FROM orders WHERE user_id = $1 AND status like '%' || $2 || '%'`
	err := db.QueryRow(queryCount, userID, status).Scan(&total)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`SELECT 
	    id, 
	    user_id, 
	    market_name, 
	    side, 
		size, 
		price,
		type,
		status,
		created_on,
		updated_on
		FROM orders WHERE user_id = $1 AND status like '%' || $2 || '%'
		OFFSET $3 LIMIT $4`, userID, status, page, pageSize)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orders := make([]*protoOrder.Order, 0)
	for rows.Next() {
		var price sql.NullFloat64
		o := new(protoOrder.Order)
		err := rows.Scan(
			&o.OrderID,
			&o.UserID,
			&o.MarketName,
			&o.Side,
			&o.Size,
			&price,
			&o.Type,
			&o.Status,
			&o.CreatedOn,
			&o.UpdatedOn,
		)

		if err != nil {
			return nil, err
		}
		if price.Valid {
			o.Price = price.Float64
		}
		orders = append(orders, o)
	}

	return &protoOrder.OrdersPage{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		Orders:   orders,
	}, nil
}

func InsertOrder(db *sql.DB, order *protoOrder.Order) error {
	sqlStatement := `insert into orders (
		id, 
		user_id, 
		market_name, 
		side, 
		size, 
		price,
		type,
		status,
		created_on, 
		updated_on) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := db.Exec(sqlStatement,
		order.OrderID,
		order.UserID,
		order.MarketName,
		order.Side,
		order.Size,
		order.Price,
		order.Type,
		order.Status,
		order.CreatedOn,
		order.UpdatedOn,
	)

	return err
}
