package entity

import (
	"gorm.io/gorm"
)

type Officer struct {
	gorm.Model
	Name     string
	Tel      string
	Email    string `gorm:"uniqueIndex"`
	Password string
}
