package storage

import (
	"gorm.io/gorm"
)

type DB struct {
	Tx *gorm.DB
}
