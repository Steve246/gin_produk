package repo

import (
	"gin_produk/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Add(newProduct *model.Product) error
	RetrieveData() []model.Product
}

type productRepository struct {
	// productDb []model.Product
	db *gorm.DB
}

func (c *productRepository) RetrieveData() []model.Product {
	var products []model.Product
	c.db.Find(&products)

	return products
}

func (p *productRepository) Add(newProduct *model.Product) error {
	err := p.db.Create(&newProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}
