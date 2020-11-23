package config

//go:generate mockgen -destination=$MOCK_DIR/store_mock.go -package=mocks -source=$GOFILE

import (
	"fmt"
	"os"
	"strings"

	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
	log "wwwin-github.cisco.com/eti/sre-go-logger"
)

const prefix = "FDS"

// configuration key constants used for configuration value lookup
const (
	ApplicationNameKey        = "FDS_APP_NAME"
	ApplicationVersionKey     = "FDS_APP_VERSION"
	MetricsURLKey             = "FDS_METRICS_URL"
	GRPCPortKey               = "FDS_GRPC_PORT"
	HTTPPortKey               = "FDS_HTTP_PORT"
	DatabaseConfigFilePathKey = "FDS_DB_CONFIG_FILE_PATH"
	OAuthConfigFilePath       = "FDS_OAUTH_CONFIG_FILE_PATH"
	SwaggerJSONURLKEY         = "FDS_SWAGGER_JSON_URL"
	InfluxDbNameKey           = "FDS_INFLUXDB_NAME"
	InfluxDbURLKey            = "FDS_INFLUXDB_URL"
	CasbinModelFilePath       = "FDS_CASBIN_MODEL_FILE_PATH"
	CasbinPolicyFilePath      = "FDS_CASBIN_POLICY_FILE_PATH"
	DatabaseConfigNotifyDir   = "FDS_DB_NOTIFY_DIR"
	DatabaseConfigNotifyFile  = "FDS_DB_NOTIFY_FILEPATH"
	MessageRouterFQDNsKey     = "FDS_MR_FQDN"
)

//Store interface for application configuration operations
type Store interface {
	Load(keymap []KeyMap)
	GetConfig(key string) (interface{}, bool)
	GetConfigString(key string) (value string, found bool)
}

//KeyMap struct holds env variables keys
type KeyMap struct {
	Name         string `json:"name"`
	UseAppPrefix bool   `json:"use_app_prefix"`
	DefaultValue string `json:"default_value"`
}

//DataServiceConfigStore struct holds app config key-value pair
type DataServiceConfigStore struct {
	log *log.Logger
	kvp map[string]interface{}
}

//NewConfigStore creates new application configuration
func NewConfigStore(log *log.Logger) Store {
	return &DataServiceConfigStore{log: log}
}

//GetConfig returns value of the key
func (cp *DataServiceConfigStore) GetConfig(key string) (interface{}, bool) {
	value, ok := cp.kvp[strings.ToUpper(key)]
	return value, ok
}

//GetConfigString returns value of the key
func (cp *DataServiceConfigStore) GetConfigString(key string) (value string, found bool) {
	var keyValue interface{}
	if keyValue, found = cp.GetConfig(key); !found || keyValue == nil {
		cp.log.Error(fmt.Sprintf("failed to get '%s' from config", key))
		return "", found
	}
	if value, found = keyValue.(string); !found || value == "" {
		cp.log.Error(fmt.Sprintf("failed to get '%s' from config", key))
		return "", found
	}
	return value, found
}

//Load takes default KeyMap kvp to as application configuration parameters.
func (cp *DataServiceConfigStore) Load(keymap []KeyMap) {
	if keymap == nil {
		keymap = getDefaultValues()
	}

	cp.kvp = make(map[string]interface{})

	for _, element := range keymap {
		key := strings.ToUpper(element.Name)

		if element.UseAppPrefix {
			key = fmt.Sprintf("%s_%s", prefix, key)
		}

		value := os.Getenv(key)

		if value == "" {
			cp.log.Info("Env key %s missing. Starting with default value %s", key, element.DefaultValue)
			value = element.DefaultValue
		}

		cp.kvp[key] = value

		cp.log.Info("config key %s has value %s", key, value)
	}
}

//This function can be separate implementation of someone else
// At that point, it can be tested separately and injected in the
// function from where it's getting called.
func getDefaultValues() []KeyMap {
	return []KeyMap{
		{
			Name:         ApplicationNameKey,
			DefaultValue: utils.DefaultAppName,
		},
		{
			Name: ApplicationVersionKey,
		},
		{
			Name: MetricsURLKey,
		},
		{
			Name:         GRPCPortKey,
			DefaultValue: utils.DefaultGRPCPort,
		},
		{
			Name:         HTTPPortKey,
			DefaultValue: utils.DefaultHTTPPort,
		},
		{
			Name: DatabaseConfigFilePathKey,
		},
		{
			Name: DatabaseConfigNotifyDir,
		},
		{
			Name: DatabaseConfigNotifyFile,
		},
		{
			Name: OAuthConfigFilePath,
		},
		{
			Name: SwaggerJSONURLKEY,
		},
		{
			Name: InfluxDbURLKey,
		},
		{
			Name:         InfluxDbNameKey,
			DefaultValue: utils.DefaultInfluxDBName,
		},
		{
			Name: CasbinModelFilePath,
		},
		{
			Name: CasbinPolicyFilePath,
		},
		{
			Name: MessageRouterFQDNsKey,
		},
	}
}

//GetConfigString for fetching env variables when config object is not loaded yet
func GetConfigString(key string) (string, bool) {
	value := os.Getenv(key)
	if value != "" {
		log.Info("Env key %s has value %s", key, value)
		return value, true
	}

	keymap := getDefaultValues()
	for _, element := range keymap {
		if element.Name == key {
			value = element.DefaultValue
			if value != "" {
				log.Info("Env key %s missing, using default: %s", key, value)
				return value, true
			}
			log.Error("Env key %s missing, no default defined", key)
			return "", false
		}
	}

	log.Error("Env key %s missing", key)
	return "", false
}
