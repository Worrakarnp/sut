package entity

import (
	"gorm.io/gorm"
)

type District struct {
	gorm.Model
	District string
	Member   []Member `gorm:"foreignKey:DistrictID"`
}
