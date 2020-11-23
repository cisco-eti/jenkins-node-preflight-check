package config

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
	log "wwwin-github.cisco.com/eti/sre-go-logger"
)

var logger *log.Logger

func init() {
	logConfig := log.DefaultConfig
	logConfig.DisableStacktrace = true
	logger, _, _ = log.New("MockService", logConfig)
}

func TestGetConfigValueExistsWhenAppPrefixIsTrue(t *testing.T) {
	cfg := NewConfigStore(logger)
	keys := []KeyMap{
		{
			Name:         "metrics_url",
			UseAppPrefix: true,
		},
	}
	expected := "http://metricsurl.cisco.com"
	assert.Nil(t, os.Setenv(MetricsURLKey, expected))
	cfg.Load(keys)

	actual, ok := cfg.GetConfig(MetricsURLKey)
	assert.Equal(t, true, ok)
	assert.Equal(t, expected, actual)

	actual, ok = cfg.GetConfig(strings.ToLower(MetricsURLKey))
	assert.Equal(t, true, ok)
	assert.Equal(t, expected, actual)
}

func TestGetConfigValueDNEWhenAppPrefixIsTrue(t *testing.T) {
	cfg := NewConfigStore(logger)
	keys := []KeyMap{
		{
			Name:         "metrics_url",
			UseAppPrefix: true,
		},
	}
	expected := "http://metricsurl.cisco.com"
	assert.Nil(t, os.Setenv(MetricsURLKey, expected))
	cfg.Load(keys)

	actual, ok := cfg.GetConfig("fds_oga_boga")
	assert.Equal(t, false, ok)
	assert.Nil(t, actual)
}

func TestGetConfigValueExistsWhenAppPrefixIsFalse(t *testing.T) {
	cfg := NewConfigStore(logger)
	keys := []KeyMap{
		{
			Name: "metrics_url",
		},
	}
	expected := "http://metricsurl.cisco.com"
	assert.Nil(t, os.Setenv("METRICS_URL", expected))
	cfg.Load(keys)

	actual, ok := cfg.GetConfig("METRICS_URL")
	assert.Equal(t, true, ok)
	assert.Equal(t, expected, actual)

	actual, ok = cfg.GetConfig("metrics_url")
	assert.Equal(t, true, ok)
	assert.Equal(t, expected, actual)

	value, ok := cfg.GetConfigString("METRICS_URL")
	assert.Equal(t, true, ok)
	assert.Equal(t, expected, value)

	value, ok = cfg.GetConfigString("UNKNOWN_URL")
	assert.Equal(t, false, ok)
	assert.Empty(t, value)
}

func TestGetConfigValueDNEWhenAppPrefixIsFalse(t *testing.T) {
	cfg := NewConfigStore(logger)
	keys := []KeyMap{
		{
			Name:         "metrics_url",
			UseAppPrefix: true,
		},
	}
	expected := "http://metricsurl.cisco.com"
	assert.Nil(t, os.Setenv(MetricsURLKey, expected))
	cfg.Load(keys)

	actual, ok := cfg.GetConfig("oga_boga")
	assert.Equal(t, false, ok)
	assert.Nil(t, actual)
}

func TestLoadsDefaultValuesWhenLoadCalledWithNilMap(t *testing.T) {
	expected := "http://metricsurl.cisco.com"
	assert.Nil(t, os.Setenv("FDS_METRICS_URL", expected))
	cfg := NewConfigStore(logger)
	cfg.Load(nil)

	actual, ok := cfg.GetConfig(MetricsURLKey)
	assert.Equal(t, true, ok)
	assert.Equal(t, expected, actual)

	actual, ok = cfg.GetConfig(strings.ToLower(MetricsURLKey))
	assert.Equal(t, true, ok)
	assert.Equal(t, expected, actual)
}

func TestGetConfigRetrievesDefaultValueWhenEnvIsNotSet(t *testing.T) {
	cfg := NewConfigStore(logger)

	cfg.Load(nil)

	actual, ok := cfg.GetConfig(InfluxDbNameKey)
	assert.Equal(t, utils.DefaultInfluxDBName, actual)
	assert.Equal(t, true, ok)

	actual, ok = cfg.GetConfig(GRPCPortKey)
	assert.Equal(t, utils.DefaultGRPCPort, actual)
	assert.Equal(t, true, ok)

	actual, ok = cfg.GetConfig(HTTPPortKey)
	assert.Equal(t, utils.DefaultHTTPPort, actual)
	assert.Equal(t, true, ok)
}

func TestGetConfigStringGlobal(t *testing.T) {

	actual, found := GetConfigString("Foo")
	assert.False(t, found)
	assert.Empty(t, actual)

	os.Setenv("Foo", "Bar")
	actual, found = GetConfigString("Foo")
	assert.True(t, found)
	assert.Equal(t, "Bar", actual)

	actual, found = GetConfigString(ApplicationNameKey)
	assert.True(t, found)
	assert.Equal(t, "sre-go-helloworld", actual)

	actual, found = GetConfigString(ApplicationVersionKey)
	assert.False(t, found)
	assert.Equal(t, "", actual)
}
