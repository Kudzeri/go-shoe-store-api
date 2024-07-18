package repositories

import (
	"github.com/Kudzeri/go-shoe-store-api/config"
	"github.com/Kudzeri/go-shoe-store-api/models"
)

func CreateUser(user *models.User) error {
	result := config.DB.Create(&user)

	return result.Error
}

func FindUserByID(id string) (models.User, error) {
	var user models.User
	result := config.DB.First(&user, "id = ?", id)
	return user, result.Error
}

func UpdateUser(id string, user *models.User) error {
	result := config.DB.Model(&user).Where("ud = ?", id).Updates(user)
	return result.Error
}

func DeleteUser(id string) error {
	result := config.DB.Delete(&models.User{}, id)
	return result.Error
}
