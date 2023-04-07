package datastore

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"

	"wwwin-github.cisco.com/eti/sre-go-sre-go-helloworld.git.git/pkg/config"
	"wwwin-github.cisco.com/eti/sre-go-sre-go-helloworld.git.git/pkg/utils"
)

func OpenDB() (*gorm.DB, error) {
	dsn, err := config.ReadDBconfig()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.Use(prometheus.New(prometheus.Config{
		DBName:          utils.DatabaseName,
		RefreshInterval: 15,
		StartServer:     false,
	}))
	if err != nil {
		return nil, err
	}

	return db, nil
}
