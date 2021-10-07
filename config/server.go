package config

var ServerConfig = map[string]interface{}{}

func InitServerConfig()  {
	ServerConfig = map[string]interface{}{
		"port" : 8080,
	}
}



