package config

import "github.com/wk331100/iFTY/system/helper"

var RedisConfig = helper.Map{}

func InitRedisConfig(){
	RedisConfig = helper.Map{
		"master" : helper.Map{
			"host" : helper.Env("redis_host", "REDIS_MASTER"),
			"port" : helper.EnvInt("redis_port", "REDIS_MASTER"),
			"password" : helper.Env("redis_password", "REDIS_MASTER"),
			"db" :  helper.EnvInt("redis_db", "REDIS_MASTER"),
		},
		"slave1" : helper.Map{
			"host" : helper.Env("redis_slave1_host", "REDIS_SLAVE1"),
			"port" : helper.EnvInt("redis_slave1_port", "REDIS_SLAVE1"),
			"password" : helper.Env("redis_slave1_password", "REDIS_SLAVE1"),
			"db" :  helper.EnvInt("redis_slave1_db", "REDIS_SLAVE1"),
		},
	}
}

