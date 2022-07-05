package manager

import (
	"gin_produk/repo"
)

type RepositoryManager interface {
	//Disini kumpulan semua repo dalam 1 projek yang dibuat

	ProductRepo() repo.ProductRepository
}

type repositoryManager struct {
	productRepo repo.ProductRepository
}

func (r *repositoryManager) ProductRepo() repo.ProductRepository {

	return r.productRepo

}

func NewRepositoryManager() RepositoryManager {
	productRepo := repo.NewProductRepository()
	return &repositoryManager{
		productRepo: productRepo,
	}
}
