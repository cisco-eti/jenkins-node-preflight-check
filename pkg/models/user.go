package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"unique"`

	IDPUserID string `gorm:"uniqueIndex:compositeidentity"`
	IDPIssuer string `gorm:"uniqueIndex:compositeidentity"`

	Sessions []Session `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

type Session struct {
	gorm.Model

	UserID      uint
	AccessToken string
	IDToken     string
}
