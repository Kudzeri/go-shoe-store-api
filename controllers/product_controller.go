package controllers

import (
	"github.com/Kudzeri/go-shoe-store-api/config"
	"github.com/Kudzeri/go-shoe-store-api/models"
)

func FindAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := config.DB.Find(&products)
	return products, result.Error
}

func FindProductByID(id string) (models.Product, error) {
	var product models.Product
	result := config.DB.First(&product, "id = ?", id)
	return product, result.Error
}

func CreateProduct(product *models.Product) error {
	result := config.DB.Create(product)
	return result.Error
}

func UpdateProduct(id string, product *models.Product) error {
	result := config.DB.Model(&product).Where("id = ?", id).Updates(product)
	return result.Error
}

func DeleteProduct(id string) error {
	result := config.DB.Delete(&models.Product{}, id)
	return result.Error
}
