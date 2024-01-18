package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Username  string         `json:"username" gorm:"unique"`
	Password  string         `json:"password" gorm:"->;<-;not null"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	IsAdmin   bool           `json:"is_admin" gorm:"default:false"`
	Metadata  string         `json:"metadata"`
}
