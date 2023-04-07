package utils

// gRPC constants
const (
	DefaultGRPCPort = "9090"
)

// http constants
const (
	OrganizationIDPathPrefix   = "orgid"
	RegionPathKey              = "regionID"
	LocationPathKey            = "locationID"
	ChannelPathKey             = "channelID"
	UserPathKey                = "userID"
	FeaturePathKey             = "feature"
	IdentityUserPathKey        = "identityUserID"
	MachineAccountIDPathPrefix = "machineAccountID"
	EntityNameKey              = "entityName"
	TrackingIDPrefix           = "sre-go-helloworld.git"
	ApplicationNameKey         = "sre-go-helloworld.git"
	DatabaseName               = "sre-go-helloworld.git"

	DefaultAppName     = "sre-go-sre-go-helloworld.git.git"
	DefaultHTTPPort    = "8080"
	ConfigTypeUser     = "user_config"
	ConfigTypeLogin    = "user_login"
	ConfigTypeLogout   = "user_logout"
	ConfigTypeLocation = "location_config"
	MetricTag          = "metric"
)

const (
	//DefaultInfluxDBName holds default DB name for Influx DB used by Data Service
	DefaultInfluxDBName = "sreInfluxDB"
)
