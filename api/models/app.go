package models

import (
	"gorm.io/gorm"
	"time"
)

type App struct {
	ID           uint           `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name         string         `json:"name" gorm:"unique"`
	IsActive     bool           `json:"is_active" gorm:"default:true"`
	RedirectURIs string         `json:"redirect_uris"` // comma separated
	ClientKey    string         `json:"client_key" gorm:"unique"`
	ClientSecret string         `json:"client_secret"`
	Scopes       string         `json:"scopes"` // comma separated
	Metadata     string         `json:"metadata"`
	AuthCode     string         `json:"auth_code"` // could be improved

	// FKs
	UserID uint `json:"user_id"`
}
