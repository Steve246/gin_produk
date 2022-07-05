package manager

import "gin_produk/repo"

// type RepositoryManager interface {
// 	//Disini kumpulan semua repo dalam 1 projek yang dibuat

// 	ProductRepo() repo.ProductRepository
// }

// type repositoryManager struct {
// 	infra infra
// }

// func (r *repositoryManager) ProductRepo() repo.ProductRepository {

// 	return repo.NewProductRepository(r.infra.SqlDb())

// }

// func NewRepositoryManager(infra infra) RepositoryManager {
// 	return &repositoryManager{
// 		infra: infra,
// 	}
// }

type RepositoryManager interface {
	// ProductRepo Disini kumpulan semua repo dalam 1 project yang dibuat
	ProductRepo() repo.ProductRepository
}

type repositoryManager struct {
	infra Infra
}

func (r *repositoryManager) ProductRepo() repo.ProductRepository {
	return repo.NewProductRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
