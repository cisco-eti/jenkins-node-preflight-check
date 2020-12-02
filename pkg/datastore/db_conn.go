package datastore

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/config"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/models"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
)

func DbConn() (db *gorm.DB, err error) {
	dsn := config.ReadDBconfig()
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	db.Use(prometheus.New(prometheus.Config{
		DBName:          utils.DatabaseName,
		RefreshInterval: 15,
		StartServer:     false,
	}))

	return db, err
}

// our initial migration function
func MigratePet() {
	db, _ := DbConn()
	// Migrate the schema
	db.AutoMigrate(&models.Pet{})
}

