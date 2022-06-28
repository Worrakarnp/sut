package entity

import (
	"gorm.io/gorm"
)

type Trader struct {
	gorm.Model
	Name    string `valid:"required~Name cannot be blank"`
	Tel     string
	Address string
}
