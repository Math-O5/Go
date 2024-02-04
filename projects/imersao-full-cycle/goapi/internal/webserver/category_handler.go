package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/Math-O5/Go/projects/imersao-full-cycle/goapi/internal/entity"
	"github.com/Math-O5/Go/projects/imersao-full-cycle/goapi/internal/service"
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
	json.NewDecoder(w).Decode(categories)
}

func (wch *WebCategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	category, _ := wch.CategoryService.GetCategory(id)

	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	json.NewDecoder(w).Encode(category)

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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewDecoder(w).Encode(result)
}
