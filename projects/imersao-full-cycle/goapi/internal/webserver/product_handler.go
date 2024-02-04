package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Math-O5/Go/projects/imersao-full-cycle/goapi/internal/entity"
	"github.com/Math-O5/Go/projects/imersao-full-cycle/goapi/internal/service"
	"github.com/go-chi/chi"
	_ "github.com/go-chi/chi"
)

type WebProductHandler struct {
	ProductService *service.ProductService
}

func NewWebProductHandler(ProductService *service.ProductService) *WebProductHandler {
	return &WebProductHandler{ProductService: ProductService}
}

func (wch *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	categories, err := wch.ProductService.GetProducts()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

func (wch *WebProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	product, _ := wch.ProductService.GetProduct(id)

	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(product)

}

func (wch *WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("SHOW ME  %s\n", product.CategoryID)

	result, err := wch.ProductService.CreateProduct(product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (wch *WebProductHandler) GetProductByCategoryID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "categoryID")
	products, _ := wch.ProductService.GetProductByCategoryID(id)

	if id == "" {
		http.Error(w, "categoryID is required", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(products)

}
