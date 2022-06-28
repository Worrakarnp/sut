package entity

import (
	"gorm.io/gorm"
)

type Degree struct {
	gorm.Model
	Degree string
	Gown   []Gown `gorm:"foreignKey:DegreeID"`
}
