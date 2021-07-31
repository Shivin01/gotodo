package models

import "gorm.io/gorm"

// Add Later on
// 1. UserPic

type User struct {
	gorm.Model
	Firstname string
	Lastname  uint
	Email     string
	Password  string
}
