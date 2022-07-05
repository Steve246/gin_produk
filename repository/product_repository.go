package repository

import "gin_produk/model"

type ProductRepo interface {
	Add(newProduct model.Product)
	RetrieveData() []model.Product
}

type productRepo struct {
	productDb []model.Product
}

func (c *productRepo) RetrieveData() []model.Product {
	return c.productDb
}

func (c *productRepo) Add(newProduct model.Product) {

	c.productDb = append(c.productDb, newProduct)

}

func NewProductRepo() ProductRepo {
	repo := new(productRepo)
	return repo
}
