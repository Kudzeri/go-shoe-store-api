package repositories

import (
	"github.com/Kudzeri/go-shoe-store-api/config"
	"github.com/Kudzeri/go-shoe-store-api/models"
)

func FindAllOrders() ([]models.Order, error) {
	var orders []models.Order
	result := config.DB.Find(&orders)
	return orders, result.Error
}

func FindOrderByID(id string) (models.Order, error) {
	var order models.Order
	result := config.DB.First(&order, id)
	return order, result.Error
}
func CreateOrder(order *models.Order) error {
	result := config.DB.Create(&order)
	return result.Error
}

func UpdateOrder(id string, order *models.Order) error {
	result := config.DB.Model(&order).Where("id = ?", id).Updates(order)
	return result.Error
}

func DeleteOrder(id string) error {
	result := config.DB.Delete(&models.Order{}, id)
	return result.Error
}
