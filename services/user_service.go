package services

import (
	"github.com/Kudzeri/go-shoe-store-api/models"
	"github.com/Kudzeri/go-shoe-store-api/repositories"
)

func RegisterUser(user *models.User) error {
	return repositories.CreateUser(user)
}

func GetUserByID(id string) (models.User, error) {
	return repositories.FindUserByID(id)
}

func UpdateUser(id string, user *models.User) error {
	return repositories.UpdateUser(id, user)
}

func DeleteUser(id string) error {
	return repositories.DeleteUser(id)
}
