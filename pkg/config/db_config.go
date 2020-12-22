package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	// DB_CONFIG_FILEPATH_ENV is an environment variable whose value is a
	// filepath to a json file that contains the attributes described by DBConfig
	DB_CONFIG_FILEPATH_ENV = "DB_CONNECTION_INFO"

	// DB_NAME_ENV is an environment variable whose value is the name of the
	// database to connect to
	DB_NAME_ENV = "DB_NAME"
)

// DBConfig holds database connection attributes
type DBConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Sslmode  string `json:"sslmode"`
	Timezone string `json:"timezone"`
}

func ReadDBconfig() (string, error) {
	filename, exist := os.LookupEnv(DB_CONFIG_FILEPATH_ENV)
	if !exist || filename == "" {
		return "", missingEnvError(DB_CONFIG_FILEPATH_ENV)
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	dbname, exist := os.LookupEnv(DB_NAME_ENV)
	if !exist || dbname == "" {
		return "", missingEnvError(DB_NAME_ENV)
	}

	var dbConfig DBConfig
	err = json.Unmarshal(file, &dbConfig)
	if err != nil {
		return "", err
	}

	connString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=%s TimeZone=%s",
		dbname,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Sslmode,
		dbConfig.Timezone)

	return connString, nil
}

func missingEnvError(envVar string) error {
	return errors.New(fmt.Sprintf("environment variable %s is not set", envVar))
}
