package models


type Product struct {
	BaseModelSoftDelete
	name string
	description string
	price float64
	stockCount uint
	sku string
	adminId uint
}

