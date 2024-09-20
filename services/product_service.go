package services

import (
	"github.com/repoleved08/ecommerce-api/config"
	"github.com/repoleved08/ecommerce-api/models"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
}

type productService struct{}

// CreateProduct implements ProductService.
func (p *productService) CreateProduct(product models.Product) (models.Product, error) {
	if err := config.DB.Create(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

// GetAllProducts implements ProductService.
func (p *productService) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func NewProductService() ProductService {
	return &productService{}
}
