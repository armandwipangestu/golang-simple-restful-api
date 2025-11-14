package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	City		string
	Street		string
	// StudentID	uint // optional one-to-one backref
}