package usecase

import (
	"gin_produk/model"
	"gin_produk/repo"
)

type CreateProductUsecase interface {
	CreateProduct(newProduct *model.Product) error
}

type createProductUsecase struct {
	repo repo.ProductRepository
}

func (c *createProductUsecase) CreateProduct(newProduct *model.Product) error {

	return c.repo.Add(newProduct)

}

func NewCreateProductUseCase(repo repo.ProductRepository) CreateProductUsecase {
	return &createProductUsecase{
		repo: repo,
	}
}
