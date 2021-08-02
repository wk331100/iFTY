package config

import (
	"github.com/wk331100/iFTY/system/helper"
)

var AppConfig = helper.Map{}

func InitConfig(){
	InitAppConfig()
	InitServerConfig()
	InitDatabaseConfig()
	InitRedisConfig()
}

func InitAppConfig()  {
	AppConfig = helper.Map{
		"pageSize" : helper.EnvInt("page_size", "APP"),
		"logPath" : helper.Env("log_path", "APP"),
		"logType" : helper.Env("log_type", "APP"),
		"timeLocation" : helper.Env("time_location", "APP"),
	}
}
