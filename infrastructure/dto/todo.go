package dto

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Text string
	Done bool
	Userid string `gorm:"type:varchar(255)"`
}