package config

import "github.com/wk331100/iFTY/system/helper"

var ServerConfig = map[string]interface{}{
	"Port" : 8080,
}

func InitServerConfig()  {
	ServerConfig = helper.Map{
		"port" : helper.EnvInt("port", "SERVER"),
	}
}