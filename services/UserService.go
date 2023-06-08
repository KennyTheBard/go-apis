package services

import "hellkite.eu/go-api/models"

func CreateUser(name string, email string) (*models.User, error) {
	user := &models.User{Name: name, Email: email}
	result := models.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
