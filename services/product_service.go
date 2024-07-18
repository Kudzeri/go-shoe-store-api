package services

import (
	"github.com/Kudzeri/go-shoe-store-api/controllers"
	"github.com/Kudzeri/go-shoe-store-api/models"
)

func GetAllProducts() ([]models.Product, error) {
	return controllers.FindAllProducts()
}

func GetProductByID(id string) (models.Product, error) {
	return controllers.FindProductByID(id)
}

func CreateProduct(product *models.Product) error {
	return controllers.CreateProduct(product)
}

func UpdateProduct(id string, product *models.Product) error {
	return controllers.UpdateProduct(id, product)
}

func DeleteProduct(id string) error {
	return controllers.DeleteProduct(id)
}
