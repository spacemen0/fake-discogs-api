package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
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

func UpdateUser(db *gorm.DB, userID uint, username, email, password string) (*User, error) {
	user, err := GetUserByID(db, userID)
	if err != nil {
		return nil, err
	}
	user.Username = username
	user.Email = email
	user.Password = password
	if err = db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(db *gorm.DB, userID uint) error {
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		return err
	}
	if err := db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
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

func GetUsersByUsername(db *gorm.DB, username string) ([]User, error) {
	var users []User

	if err := db.Where("username LIKE ?", "%"+username+"%").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func UsernameExist(db *gorm.DB, username string) bool {
	var count int64
	if err := db.Model(&User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func EmailExist(db *gorm.DB, email string) bool {
	var count int64
	if err := db.Model(&User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}
