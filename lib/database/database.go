package database

import (
	"github.com/darienkentanu/test-erajaya/models"
	"gorm.io/gorm"
)

type ProductDB struct {
	db *gorm.DB
}

func NewProductDB(db *gorm.DB) *ProductDB {
	return &ProductDB{db: db}
}

type ProductModel interface {
	AddProduct(product models.Product) (models.Product, error)
	GetNewestProduct() ([]models.Product, error)
	GetProductASC() ([]models.Product, error)
	GetProductDESC() ([]models.Product, error)
	GetProductLowestPrice() ([]models.Product, error)
	GetProductHighestPrice() ([]models.Product, error)
}

func (pdb *ProductDB) AddProduct(product models.Product) (models.Product, error) {
	if err := pdb.db.Save(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (pdb *ProductDB) GetNewestProduct() ([]models.Product, error) {
	var products []models.Product
	if err := pdb.db.Order("created_at desc").Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func (pdb *ProductDB) GetProductASC() ([]models.Product, error) {
	var products []models.Product
	if err := pdb.db.Order("name asc").Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func (pdb *ProductDB) GetProductDESC() ([]models.Product, error) {
	var products []models.Product
	if err := pdb.db.Order("name desc").Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func (pdb *ProductDB) GetProductLowestPrice() ([]models.Product, error) {
	var products []models.Product
	if err := pdb.db.Order("price asc").Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func (pdb *ProductDB) GetProductHighestPrice() ([]models.Product, error) {
	var products []models.Product
	if err := pdb.db.Order("price desc").Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}
