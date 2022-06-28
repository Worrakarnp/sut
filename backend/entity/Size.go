package entity

import (
	"gorm.io/gorm"
)

type Size struct {
	gorm.Model
	Size string
	Gown []Gown `gorm:"foreignKey:SizeID"`
}
