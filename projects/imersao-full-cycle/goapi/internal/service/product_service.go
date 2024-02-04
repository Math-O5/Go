package service

import (
	datab "github.com/Math-O5/Go/projects/imersao-full-cycle/goapi/internal/database"
	"github.com/Math-O5/Go/projects/imersao-full-cycle/goapi/internal/entity"
)

type ProductService struct {
	ProductDB datab.ProductDB
}

func NewProductService(cat datab.ProductDB) *ProductService {
	return &ProductService{ProductDB: cat}
}
func (pd *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := pd.ProductDB.GetProducts()

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (pd *ProductService) CreateCategory(name string, description string, price float64, categoryID string, imageURL string) (*entity.Product, error) {
	product := entity.NewProduct(name, description, price, categoryID, imageURL)

	_, err := pd.ProductDB.Createproducts(product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pd *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := pd.ProductDB.GetProduct(id)

	if err != nil {
		return nil, err
	}

	return product, nil

}
