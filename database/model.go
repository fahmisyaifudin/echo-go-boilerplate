package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	AuthKey   string `gorm:"unique;not null"`
}
