package config

import "github.com/wk331100/iFTY/system/helper"

var RedisConfig = helper.Map{}

func InitRedisConfig(){
	RedisConfig = helper.Map{
		"master" : helper.Map{
			"host" : helper.Env("host", "REDIS_MASTER"),
			"port" : helper.EnvInt("port", "REDIS_MASTER"),
			"password" : helper.Env("password", "REDIS_MASTER"),
			"db" :  helper.EnvInt("db", "REDIS_MASTER"),
		},
		"slave1" : helper.Map{
			"host" : helper.Env("host", "REDIS_SLAVE1"),
			"port" : helper.EnvInt("port", "REDIS_SLAVE1"),
			"password" : helper.Env("password", "REDIS_SLAVE1"),
			"db" :  helper.EnvInt("db", "REDIS_SLAVE1"),
		},
	}
}

