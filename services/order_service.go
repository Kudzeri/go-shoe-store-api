package services

import (
	"github.com/Kudzeri/go-shoe-store-api/models"
	"github.com/Kudzeri/go-shoe-store-api/repositories"
)

func GetAllOrders() ([]models.Order, error) {
	return repositories.FindAllOrders()
}

func GetOrderById(id string) (models.Order, error) {
	return repositories.FindOrderByID(id)
}

func CreateOrder(order *models.Order) error {
	return repositories.CreateOrder(order)
}

func UpdateOrder(id string, product *models.Order) error {
	return repositories.UpdateOrder(id, product)
}

func DeleteOrder(id string) error {
	return repositories.DeleteOrder(id)
}
