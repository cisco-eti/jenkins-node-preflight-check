package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

//DBConfig struct holds database connection attributes
type DBConfig struct {
	Dbname   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Sslmode  string `json:"sslmode"`
	Timezone string `json:"timezone"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadDBconfig() string {
	filename, exits := os.LookupEnv("DB_CONNECTION_INFO")
	if exits != true {
		panic(errors.New("Environment variable DB_CONNECTION_INFO is not set"))
	}
	file, err := ioutil.ReadFile(filename)

	// we initialize our Users array
	var dbConfig DBConfig

	// we unmarshal our byteArray which contains our
	// file content into 'dbConfig' which we defined above
	err = json.Unmarshal(file, &dbConfig)
	check(err)
	connString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=%s TimeZone=%s",
		dbConfig.Dbname,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Sslmode,
		dbConfig.Timezone)

	return connString
}
