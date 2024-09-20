package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/repoleved08/ecommerce-api/models"
	"github.com/repoleved08/ecommerce-api/services"
)

type ProductHandler struct {
	ProductService services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: service,
	}
}

func (handler *ProductHandler) CreateProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return sendErrorResponse(c, http.StatusBadRequest, "Invalid input!")
	}
	// logic for only admins creating products
	createProduct, err := handler.ProductService.CreateProduct(product)
	if err != nil {
		return sendErrorResponse(c, http.StatusInternalServerError, "an error occured while creating product")
	}
	return c.JSON(http.StatusCreated, createProduct)

}
