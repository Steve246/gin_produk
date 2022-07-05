package repo

import "gin_produk/model"

type ProductRepository interface {
	Add(newProduct *model.Product) error
	RetrieveData() []model.Product
}

type productRepository struct {
	productDb []model.Product
}

func (c *productRepository) RetrieveData() []model.Product {
	return c.productDb
}

func (p *productRepository) Add(newProduct *model.Product) error {
	p.productDb = append(p.productDb, *newProduct)
	return nil
}

func NewProductRepository() ProductRepository {
	repo := new(productRepository)
	return repo
}
