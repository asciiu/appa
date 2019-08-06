package sql

import (
	"database/sql"

	"github.com/asciiu/appa/order-service/models"
)

func DeleteOrder(db *sql.DB, orderID, userID string) error {
	_, err := db.Exec("DELETE FROM orders WHERE id = $1 and user_id = $2", orderID, userID)
	return err
}

func FindOrderByID(db *sql.DB, orderID string) (*models.Order, error) {
	o := new(models.Order)
	err := db.QueryRow(`SELECT 
	    id, 
	    user_id, 
	    market_name, 
	    side, 
		amount, 
		filled,
		price,
		type,
		status,
		created_on,
		updated_on
	    FROM orders WHERE id = $1`, orderID).Scan(
		&o.ID,
		&o.UserID,
		&o.MarketName,
		&o.Side,
		&o.Amount,
		&o.Filled,
		&o.Price,
		&o.Type,
		&o.Status,
		&o.CreatedOn,
		&o.UpdatedOn,
	)

	if err != nil {
		return nil, err
	}
	return o, nil
}

func FindUserOrders(db *sql.DB, userID, status string, page, pageSize uint32) (*models.OrdersPage, error) {
	pagedOrders := &models.OrdersPage{
		Page:     page,
		PageSize: pageSize,
		Orders:   make([]*models.Order, 0),
	}

	queryCount := `SELECT count(*) FROM orders 
	               WHERE user_id = $1 AND status like '%' || $2 || '%'`
	err := db.QueryRow(queryCount, userID, status).Scan(&pagedOrders.Total)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`SELECT 
	    id, 
	    user_id, 
	    market_name, 
	    side, 
		amount, 
		filled,
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

	for rows.Next() {
		var price sql.NullInt64
		o := new(models.Order)
		err := rows.Scan(
			&o.ID,
			&o.UserID,
			&o.MarketName,
			&o.Side,
			&o.Amount,
			&o.Filled,
			&price,
			&o.Type,
			&o.Status,
			&o.CreatedOn,
			&o.UpdatedOn,
		)

		if err != nil {
			return nil, err
		}
		pagedOrders.Orders = append(pagedOrders.Orders, o)
	}

	return pagedOrders, nil
}

func InsertOrder(db *sql.DB, order *models.Order) error {
	sqlStatement := `insert into orders (
		id, 
		user_id, 
		market_name, 
		side, 
		amount, 
		filled,
		price,
		type,
		status,
		created_on, 
		updated_on) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := db.Exec(sqlStatement,
		order.ID,
		order.UserID,
		order.MarketName,
		order.Side,
		order.Amount,
		order.Filled,
		order.Price,
		order.Type,
		order.Status,
		order.CreatedOn,
		order.UpdatedOn,
	)

	return err
}

func UpdateOrderStatus(db *sql.DB, orderID, status string) error {
	_, err := db.Exec("UPDATE orders SET status = $1 WHERE id = $2",
		status, orderID)

	return err
}

func UpdateOrderAmounts(db *sql.DB, orderID string, amount, filled uint64) error {
	_, err := db.Exec("UPDATE orders SET amount = $1, filled = $2 WHERE id = $3",
		amount, filled, orderID)

	return err
}
