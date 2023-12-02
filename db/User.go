package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email                 string
	Password              string
	UserTypeID            uint16
	EmailConfirmed        bool
	EmailConfirmationCode string
}
