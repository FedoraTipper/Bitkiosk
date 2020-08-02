package product

import (
	"github.com/jinzhu/gorm"
)

func LoadActiveProducts(db *gorm.DB) []Product {
	var activeProducts []Product

	db.Where("start_date <= now() and (end_date is NULL or (start_date != end_date and end_date > now()))").Find(&activeProducts)

	return activeProducts
}

func LoadProductWithId(id uint, db *gorm.DB) Product {
	var product Product

	db.Where("id = ?", id).First(&product)

	return product
}

func LoadProductFromSku(sku string, db *gorm.DB) Product {
	var product Product

	db.Where("sku = ?", sku).First(&product)

	return product
}