package usecase

import (
	"gin_produk/model"
	"gin_produk/repo"
)

type AllProductUsecase interface {
	displayAll() []model.Product
}

type allProductUsecase struct {
	repo repo.ProductRepository
}

func (c *allProductUsecase) displayAll() []model.Product {

	alldata := c.repo.RetrieveData()

	return alldata
}

func NewAllProductUsecase(repo repo.ProductRepository) AllProductUsecase {
	return &allProductUsecase{
		repo: repo,
	}
}
