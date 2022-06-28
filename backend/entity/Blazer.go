package entity

import (
	"gorm.io/gorm"
)

type Blazer struct {
	gorm.Model
	Blazer string
	Gown   []Gown `gorm:"foreignKey:BlazerID"`
}
