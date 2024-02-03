package datab

import (
	"database/sql"

	"github.com/Math-O5/Go/projects/imersao-full-cycle/goapi/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func newProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (cd *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := cd.db.Query("SELECT id, name FROM products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var Product entity.Product
		if err := rows.Scan(&Product.ID, &Product.Name); err != nil {
			return nil, err
		}

		products = append(products, &Product)
	}

	return products, nil
}

func (cd *ProductDB) Createproducts(product *entity.Product) (string, error) {
	_, err := cd.db.Exec("INSERT INTO products (id, name, description, price, category_id, image_url) VALUES (?,?,?,?,?)", product.ID, product.Name)

	if err != nil {
		return "", err
	}

	return product.ID, nil
}

func (cd *ProductDB) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product

	err := cd.db.QueryRow("SELECT id, name, price, category_id, image_url FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.ImageURL)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (pd *ProductDB) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT id, name FROM products WHERE category_id = ?", categoryID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var Product entity.Product
		if err := rows.Scan(&Product.ID, &Product.Name); err != nil {
			return nil, err
		}

		products = append(products, &Product)
	}

	return products, nil
}
