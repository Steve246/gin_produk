package repo

import (
	"gin_produk/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Add(newProduct *model.Product) error
	Retrieve() ([]model.Product, error)
}

type productRepository struct {
	// productDb []model.Product
	db *gorm.DB
}

func (c *productRepository) Retrieve() ([]model.Product, error) {
	var products []model.Product
	err := c.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
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
