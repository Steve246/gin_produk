package usecase

import (
	"gin_produk/model"
	"gin_produk/repo"
)

type AllProductUsecase interface {
	DisplayAll() []model.Product
}

type allProductUsecase struct {
	repo repo.ProductRepository
}

func (c *allProductUsecase) DisplayAll() []model.Product {

	alldata := c.repo.RetrieveData()

	return alldata
}

func NewAllProductUsecase(repo repo.ProductRepository) AllProductUsecase {
	return &allProductUsecase{
		repo: repo,
	}
}
