package storage

import (
	"github.com/sirrobot01/oauth-sso/api/models"
)

func (db *DB) GetApps() (apps []models.App, err error) {
	err = db.Tx.Find(&apps).Error
	return
}

func (db *DB) GetAppById(id string) (app *models.App, err error) {
	err = db.Tx.First(&app, id).Error
	return
}

func (db *DB) GetAppByQuery(query string, args ...any) (app *models.App, err error) {
	// SELECT * FROM apps WHERE id = ? LIMIT 1
	err = db.Tx.Where(query, args).First(&app).Error
	return
}

func (db *DB) CreateApp(app *models.App) (err error) {
	err = db.Tx.Create(&app).Error
	return
}
