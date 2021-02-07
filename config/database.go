package config

import "github.com/wk331100/iFTY/system/helper"

var MysqlConfig = helper.Map{}
var RedisConfig = helper.Map{}

func InitConfig(){
	MysqlConfig = helper.Map{
		"master" : helper.Map{
			"host" : helper.Env("host", "DB_MASTER"),
			"port" : helper.EnvInt("port", "DB_MASTER"),
			"username" : helper.Env("username", "DB_MASTER"),
			"password" : helper.Env("password", "DB_MASTER"),
			"dbname" : helper.Env("dbname", "DB_MASTER"),
		},
		"slave" : helper.Map{
			"host" : helper.Env("host", "DB_SLAVE"),
			"port" : helper.EnvInt("port", "DB_SLAVE"),
			"username" : helper.Env("username", "DB_SLAVE"),
			"password" : helper.Env("password", "DB_SLAVE"),
			"dbname" : helper.Env("dbname", "DB_SLAVE"),
		},
	}

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

