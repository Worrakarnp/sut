package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Gown struct {
	gorm.Model
	Name    string `valid:"required~Name cannot be blank"`
	Student string
	Status  string
	Tel     string
	Date    time.Time `valid:"notFuture~Date cannot be in the future"`

	//Size ทำหน้าที่เป็น FK
	SizeID *uint
	Size   Size `gorm:"references:id" valid:"-"`

	//Blazer ทำหน้าที่เป็น FK
	BlazerID *uint
	Blazer   Blazer `gorm:"references:id" valid:"-"`

	//Degree ทำหน้าที่เป็น FK
	DegreeID *uint
	Degree   Degree `gorm:"references:id" valid:"-"`

	//Branch ทำหน้าที่เป็น FK
	BranchID *uint
	Branch   Branch `gorm:"references:id" valid:"-"`
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
