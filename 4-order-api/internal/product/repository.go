package product

import (
	"ps-go-adv/4-order-api/pkg/db"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(database *db.Db) *ProductRepository {
	return &ProductRepository{
		Database: database,
	}
}

func (repo *ProductRepository) Create(product *Product) (*Product, error) {
	result := repo.Database.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) Update(product *Product) (*Product, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) Delete(id uint) (*gorm.DB, error) {
	result := repo.Database.DB.Delete(&Product{}, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (repo *ProductRepository) GetById(id uint) (*Product, error) {
	var product Product
	result := repo.Database.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repo *ProductRepository) GetAll() ([]Product, error) {
	var products []Product
	result := repo.Database.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}