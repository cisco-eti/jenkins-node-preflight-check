package config

//DBConfig struct holds database connection attributes
type DBConfig struct {
	DBStatus      string `json:"dbstatus"`
	Type          string `json:"type"`
	Host          string `json:"host"`
	Port          int    `json:"port"`
	User          string `json:"user"`
	Password      string `json:"-"`
	DBName        string `json:"dbname"`
	SSLMode       string `json:"sslmode"`
	SchemaVersion int    `json:"dbschemaversion"`
}
