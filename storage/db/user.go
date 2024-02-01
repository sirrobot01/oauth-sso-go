package storage

import (
	"github.com/sirrobot01/oauth-sso/api/models"
)

func (db *DB) GetUsers() (users []models.User, err error) {
	err = db.Tx.Find(&users).Error
	return
}

func (db *DB) GetUserById(id string) (user *models.User, err error) {
	err = db.Tx.First(&user, id).Error
	return
}

func (db *DB) GetUserByQuery(query string, args ...any) (user *models.User, err error) {
	// SELECT * FROM users WHERE username = ? LIMIT 1
	err = db.Tx.Where(query, args).First(&user).Error
	return
}

func (db *DB) CreateUser(user *models.User) (err error) {
	err = db.Tx.Create(&user).Error
	return
}
