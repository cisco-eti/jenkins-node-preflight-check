package utils

import (
	etilog "wwwin-github.cisco.com/eti/sre-go-logger"
)

func Loginit() *etilog.Logger {
	//Fetch service config
	appName := ApplicationNameKey

	//Initialize Logger
	logConfig := etilog.DefaultProdConfig
	logConfig.DisableStacktrace = true
	logger, stop, err := etilog.New(appName, logConfig)
	if err != nil {
		logger, stop = etilog.Default(appName)
		logger.Error("*** FAILED TO CREATE ETI LOGGER %+v", err)
	}
	defer stop()
	//Set Logging Prefix for Data Service
	logger.SetTrackingIDPrefix(TrackingIDPrefix)
	return logger
}
