package entity

import (
	"gorm.io/gorm"
)

type Branch struct {
	gorm.Model
	Branch    string
	Member    []Member    `gorm:"foreignKey:BranchID"`
	Syndicate []Syndicate `gorm:"foreignKey:BranchID"`
}
