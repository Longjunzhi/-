package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"Name" gorm:"type:varchar(100);unique_index;not null;default ''"`
	Password string `json:"-" gorm:"size:255;index;not null"`
	Sex      int8   `json:"sex" gorm:"not null;default 0"`
}

func NewUser() (user *User) {
	return &User{}
}
