package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	ID        string `gorm:"default:uuid_generate_v3()"`
	Name      string `gorm:"not null;type:varchar(100)"`
	Email     string `gorm:"unique;not null;type:varchar(100);default:null"`
	Password  string `gorm:"not null" `
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
