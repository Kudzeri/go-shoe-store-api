package services

import (
	"github.com/Kudzeri/go-shoe-store-api/models"
	"github.com/Kudzeri/go-shoe-store-api/repositories"
)

func GetAllProducts() ([]models.Product, error) {
	return repositories.FindAllProducts()
}

func GetProductByID(id string) (models.Product, error) {
	return repositories.FindProductByID(id)
}

func CreateProduct(product *models.Product) error {
	return repositories.CreateProduct(product)
}

func UpdateProduct(id string, product *models.Product) error {
	return repositories.UpdateProduct(id, product)
}

func DeleteProduct(id string) error {
	return repositories.DeleteProduct(id)
}
