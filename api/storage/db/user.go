package db

import (
	"github.com/sirrobot01/oauth-sso/api/models"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) (users []models.User, err error) {
	err = db.Find(&users).Error
	return
}

func GetUserById(db *gorm.DB, id int) (user models.User, err error) {
	err = db.First(&user, id).Error
	return
}

func GetUserByUsername(db *gorm.DB, username string) (user models.User, err error) {
	// SELECT * FROM users WHERE username = ? LIMIT 1
	err = db.Where("username = ?", username).First(&user).Error
	return
}

func CreateUser(db *gorm.DB, user *models.User) (err error) {
	err = db.Create(&user).Error
	return
}
