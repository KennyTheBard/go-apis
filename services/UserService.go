package services

import "hellkite.eu/go-api/models"

func CreateUser(name string, email string) (models.User, error) {
	user := models.User{Name: name, Email: email}
	if err := models.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User

	if err := models.DB.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func GetUserById(id int) (models.User, error) {
	var user models.User
	err := models.DB.First(&user, id).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUserById(id int, name string) (models.User, error) {
	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	user.Name = name
	if err := models.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
