package orders

import (
	"project/internal/database"
)

type Repository interface {
	Create(order Order) error
	Cancel(order_number string, user_id int) error
	Detail(order_number string, user_id int) (Order, error)
	FindAll() ([]Order, error)
	FindOrderByUserID(user_id int) ([]Order, error)
}

type orderRepository struct {
	db *database.PostgresDB
}

func NewOrderRepository(db *database.PostgresDB) Repository {
	return &orderRepository{
		db: db,
	}
}

func FindOrderItems(order_id int, r *orderRepository) (*[]OrderItem, error) {
	var order_items []OrderItem
	query := "SELECT id, order_id, product_id, quantity, price, created_at, updated_at FROM order_items WHERE order_id = $1"
	rows, err := r.db.Conn.Query(query, order_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order_item OrderItem
		err := rows.Scan(&order_item.ID,
			&order_item.OrderID,
			&order_item.ProductID,
			&order_item.Quantity,
			&order_item.Price,
			&order_item.CreatedAt,
			&order_item.UpdatedAt)
		if err != nil {
			return nil, err
		}
		order_items = append(order_items, order_item)
	}

	return &order_items, nil
}

func (r *orderRepository) Create(order Order) error {
	query := "INSERT INTO orders (order_number, status, shipping_address, user_id, total, total_discount) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	err := r.db.Conn.QueryRow(query,
		order.OrderNumber,
		order.Status,
		order.ShippingAddress,
		order.UserID,
		order.Total,
		order.TotalDiscount).Scan(&order.ID)
	if err != nil {
		return err
	}

	for _, item := range *order.OrderItems {
		query := "INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4) RETURNING id"
		err = r.db.Conn.QueryRow(query,
			order.ID,
			item.ProductID,
			item.Quantity,
			item.Price).Scan(&item.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *orderRepository) Cancel(order_number string, user_id int) error {
	query := "UPDATE orders SET is_cancelled = TRUE, updated_at = current_timestamp WHERE order_number = $1 AND user_id = $2"
	_, err := r.db.Conn.Exec(query, order_number, user_id)
	if err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) Detail(order_number string, user_id int) (Order, error) {
	var order Order
	query := "SELECT id, order_number, status, shipping_address, user_id, total, total_discount, is_cancelled, created_at, updated_at FROM orders WHERE order_number = $1 AND user_id = $2"
	err := r.db.Conn.QueryRow(query, order_number, user_id).Scan(
		&order.ID,
		&order.OrderNumber,
		&order.Status,
		&order.ShippingAddress,
		&order.UserID,
		&order.Total,
		&order.TotalDiscount,
		&order.IsCancelled,
		&order.CreatedAt,
		&order.UpdatedAt)
	if err != nil {
		return order, err
	}
	order_items, _ := FindOrderItems(order.ID, r)
	order.OrderItems = order_items

	return order, nil
}

func (r *orderRepository) FindAll() ([]Order, error) {
	var orders []Order
	query := "SELECT id, order_number, status, shipping_address, user_id, total, total_discount FROM orders WHERE is_cancelled IS FALSE"
	rows, err := r.db.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		err := rows.Scan(&order.OrderNumber,
			&order.Status,
			&order.ShippingAddress,
			&order.UserID,
			&order.Total,
			&order.TotalDiscount)
		if err != nil {
			return nil, err
		}
		order_items, _ := FindOrderItems(order.ID, r)
		order.OrderItems = order_items
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *orderRepository) FindOrderByUserID(user_id int) ([]Order, error) {
	var orders []Order
	query := "SELECT id, order_number, status, shipping_address, user_id, total, total_discount FROM orders WHERE user_id = $1"
	rows, err := r.db.Conn.Query(query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		err := rows.Scan(
			&order.ID,
			&order.OrderNumber,
			&order.Status,
			&order.ShippingAddress,
			&order.UserID,
			&order.Total,
			&order.TotalDiscount)
		if err != nil {
			return nil, err
		}
		order_items, _ := FindOrderItems(order.ID, r)
		order.OrderItems = order_items
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}
