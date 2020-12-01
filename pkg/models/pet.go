package models

import "gorm.io/gorm"

// Our User Struct
type Pet struct {
	gorm.Model
	Name   string
	Type   string
	Family string
}

