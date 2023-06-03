package products

import "project/internal/database"

type Repository interface {
	FindAll() (*[]Product, error)
	FindByCode(code string) (*Product, error)
	Create(product Product) error
	FindID(code string) (*int, error)
}

type productRepository struct {
	db *database.PostgresDB
}

func NewProductRepository(db *database.PostgresDB) Repository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) FindAll() (*[]Product, error) {
	var products []Product
	query := "SELECT * FROM products"
	rows, err := r.db.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Code, &product.ProductName, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return &products, nil
}

func (r *productRepository) FindByCode(code string) (*Product, error) {
	var product Product
	query := "SELECT * FROM products WHERE code = $1"
	err := r.db.Conn.QueryRow(query, code).Scan(
		&product.ID, &product.Code, &product.ProductName, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Create(product Product) error {
	query := "INSERT INTO products (code, product_name, price) VALUES ($1, $2, $3)"
	_, err := r.db.Conn.Exec(query,
		product.Code,
		product.ProductName,
		product.Price)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) FindID(code string) (*int, error) {
	var product Product
	query := "SELECT id FROM products WHERE code = $1"
	err := r.db.Conn.QueryRow(query, code).Scan(
		&product.ID)
	if err != nil {
		return nil, err
	}
	return &product.ID, nil
}
