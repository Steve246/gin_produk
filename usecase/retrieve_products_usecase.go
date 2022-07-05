package usecase

import (
	"gin_produk/model"
	"gin_produk/repo"
)

type AllProductUsecase interface {
	DisplayAll() ([]model.Product, error)
}

type allProductUsecase struct {
	repo repo.ProductRepository
}

func (c *allProductUsecase) DisplayAll() ([]model.Product, error) {
	return c.repo.Retrieve()
}

func NewAllProductUsecase(repo repo.ProductRepository) AllProductUsecase {
	return &allProductUsecase{
		repo: repo,
	}
}
