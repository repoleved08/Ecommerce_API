package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model  `swaggerignore:"true"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	ImageURL    string  `json:"image_url"`
}
