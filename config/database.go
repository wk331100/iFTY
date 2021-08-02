package config

import "github.com/wk331100/iFTY/system/helper"

var MysqlConfig = helper.Map{}

func InitDatabaseConfig(){
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
}

