package models

import (
	"errors"
	"fake-discogs-api/config"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"unique;not null" json:"username" binding:"required"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password" binding:"required"`
	Bio      string `json:"bio"`
}

func CreateUser(db *gorm.DB, username, email, password string) (*User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Username = username
	user.Email = email
	user.Password = string(hashedPassword)
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
	if err := db.Select("id, username, email, bio").First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUsernameByID(db *gorm.DB, userID uint) (string, error) {
	var user User
	if err := db.Select("username").First(&user, userID).Error; err != nil {
		return "", err
	}

	return user.Username, nil
}

func GetUserByUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	if err := db.Select("username, bio").Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUsersByUsername(db *gorm.DB, username string) ([]User, error) {
	var users []User

	if err := db.Select("username, bio").Where("username LIKE ?", "%"+username+"%").Find(&users).Error; err != nil {
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

func UserLoginByUsername(db *gorm.DB, username, password string) (string, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := generateJWTToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func UserLoginByEmail(db *gorm.DB, email, password string) (string, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := generateJWTToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateJWTToken(user_id uint) (string, error) {
	c := config.GetConfig()

	claims := jwt.MapClaims{
		"user_id": user_id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(c.GetString("jwt.secret"))
	tokenString, err := token.SignedString(secretKey)
	return tokenString, err
}
