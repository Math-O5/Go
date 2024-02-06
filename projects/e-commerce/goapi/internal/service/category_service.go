package service

import (
	datab "github.com/Math-O5/Go/projects/imersao-full-cycle/goapi/internal/database"
	"github.com/Math-O5/Go/projects/imersao-full-cycle/goapi/internal/entity"
)

type CategoryService struct {
	CategoryDB datab.CategoryDB
}

func NewCategoryService(cat datab.CategoryDB) *CategoryService {
	return &CategoryService{CategoryDB: cat}
}
func (cs *CategoryService) GetCategories() ([]*entity.Category, error) {
	categories, err := cs.CategoryDB.GetCategories()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (cs *CategoryService) CreateCategory(name string) (*entity.Category, error) {
	category := entity.NewCategory(name)

	_, err := cs.CategoryDB.CreateCategories(category)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (cs *CategoryService) GetCategory(id string) (*entity.Category, error) {
	category, err := cs.CategoryDB.GetCategory(id)

	if err != nil {
		return nil, err
	}

	return category, nil

}
