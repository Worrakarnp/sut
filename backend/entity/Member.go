package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Name    string `valid:"required~Name cannot be blank"`
	Num     string
	Email   string
	Tel     string
	Address string
	Date    time.Time `valid:"notFuture~Date cannot be in the future"`

	//Branch ทำหน้าที่เป็น FK
	BranchID *uint
	Branch   Branch `gorm:"references:id" valid:"-"`

	//Category ทำหน้าที่เป็น FK
	CategoryID *uint
	Category   Category `gorm:"references:id" valid:"-"`

	//SubDistrict ทำหน้าที่เป็น FK
	SubDistrictID *uint
	SubDistrict   SubDistrict `gorm:"references:id" valid:"-"`

	//District ทำหน้าที่เป็น FK
	DistrictID *uint
	District   District `gorm:"references:id" valid:"-"`

	//Province ทำหน้าที่เป็น FK
	ProvinceID *uint
	Province   Province `gorm:"references:id" valid:"-"`
}

// ตรวจสอบวันที่ต้องไม่เป็น อนาคต
func init() {
	govalidator.CustomTypeTagMap.Set("positive", func(i interface{}, context interface{}) bool {
		num := i
		return num.(int) >= 1
	})
	govalidator.CustomTypeTagMap.Set("notFuture", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		now := time.Now()
		return t.Before(now) || t.Equal(now)
	})
}
