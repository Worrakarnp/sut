package entity

import (
	"gorm.io/gorm"
)

type SubDistrict struct {
	gorm.Model
	SubDistrict string
	Member      []Member `gorm:"foreignKey:SubDistrictID"`
}
