package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name		string		`gorm:"size:100;not null"`
	Age			int			`gorm:"not null"`
	Address		Address		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AddressID	*uint
}