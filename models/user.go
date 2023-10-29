package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func CreateUser(db *gorm.DB, username, email, password string) (*User, error) {
	user := &User{
		Username: username,
		Email:    email,
		Password: password,
	}

	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByID(db *gorm.DB, userID uint) (*User, error) {
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
