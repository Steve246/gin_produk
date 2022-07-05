package repository

import (
	"gin_produk/model"
)

type CategoryRepo interface {
	Add(newCategory *model.Category) ([]model.Category, error)
	Retrieve() []model.Category
}

type categoryRepo struct {
	categoryDb []model.Category
}

func (c *categoryRepo) Retrieve() []model.Category {

	// for index, isi := range modelCategory {
	// 	fmt.Println(index, isi.CategoryId, isi.CategoryName)
	// }

	return c.categoryDb

}

func (c *categoryRepo) Add(newCategory *model.Category) ([]model.Category, error) {

	addCategory := append(c.categoryDb, *newCategory)
	return addCategory, nil

}

func NewCategoryRepo() CategoryRepo {
	repo := new(categoryRepo)
	return repo
}
