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
	// 	var products []model.Product
	// 	// err := c.db.Find(&products)
	// 	c.db.Find(&products)
	// 	// if err != nil {
	// 	// 	return nil, err
	// 	// }
	// 	// return products, nil
	// 	return products
	// }

	panic("error")
}

func (p *productRepository) Add(newProduct *model.Product) error {
	// err := p.db.Create(&newProduct).Error
	// if err != nil {
	// 	return err

	// }
	// return nil
	panic("error")
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}
