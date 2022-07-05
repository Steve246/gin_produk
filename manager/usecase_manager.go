package manager

import "gin_produk/usecase"

type UseCaseManager interface {
	CreateProductUseCase() usecase.CreateProductUsecase
	ListProductUseCase() usecase.AllProductUsecase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) CreateProductUseCase() usecase.CreateProductUsecase {
	return usecase.NewCreateProductUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) ListProductUseCase() usecase.AllProductUsecase {
	return usecase.NewAllProductUsecase(u.repoManager.ProductRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
