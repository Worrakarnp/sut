package entity

import (
	"gorm.io/gorm"
)

type Syndicate struct {
	gorm.Model
	Name string `valid:"minstringlength(5)~Name must be more than 5"`
	Rank string `valid:"required~Rank cannot be blank"`
	Tel  string

	//Branch ทำหน้าที่เป็น FK
	BranchID *uint
	Branch   Branch `gorm:"references:id" valid:"-"`
}
