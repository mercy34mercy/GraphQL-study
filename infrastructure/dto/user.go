package dto

import "gorm.io/gorm"

type User struct{
	gorm.Model
	Name string `gorm:"type:varchar(255)"`
}