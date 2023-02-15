package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string `valid:"required~Name cannot be blank"`
	Email      string
	CustomerID string `valid:"matches(^[LMH]\\d{7}$)"`
}
