package entity

import (
	"gorm.io/gorm"
)

type Province struct {
	gorm.Model
	Province string
	Member   []Member `gorm:"foreignKey:ProvinceID"`
}
