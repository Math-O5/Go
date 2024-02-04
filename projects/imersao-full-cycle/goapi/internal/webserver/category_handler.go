package webserver

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Math-O5/Go/projects/imersao-full-cycle/goapi/internal/entity"
	"github.com/Math-O5/Go/projects/imersao-full-cycle/goapi/internal/service"
	"github.com/go-chi/chi"
	_ "github.com/go-chi/chi"
)

type WebCategoryHandler struct {
	CategoryService *service.CategoryService
}

func NewWebCategoryHandler(categoryService *service.CategoryService) *WebCategoryHandler {
	return &WebCategoryHandler{CategoryService: categoryService}
}

func (wch *WebCategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := wch.CategoryService.GetCategories()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

func (wch *WebCategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	category, _ := wch.CategoryService.GetCategory(id)
	var w1 io.Reader

	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	json.NewDecoder(w1).Decode(category)

}

func (wch *WebCategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.Category

	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := wch.CategoryService.CreateCategory(category.Name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
}
