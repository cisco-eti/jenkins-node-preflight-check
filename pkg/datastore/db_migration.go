package datastore

import (
	"gorm.io/gorm"

	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/models"
)

func Migrate(db *gorm.DB) error {
	err := migratePet(db)
	if err != nil {
		return err
	}

	return nil
}

func migratePet(db *gorm.DB) error {
	return db.AutoMigrate(&models.Pet{})
}
