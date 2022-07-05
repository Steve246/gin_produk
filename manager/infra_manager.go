package manager

import (
	"gin_produk/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Infra interface {
	SqlDb() *gorm.DB
}

type infra struct {
	db *gorm.DB
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func NewInfra(config *config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName)

	if err != nil {
		log.Fatal(err.Error())
	}
	return &infra{db: resource}
}

func initDbResource(DataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(DataSourceName), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil

}
