package config

import "github.com/wk331100/iFTY/system/helper"

var MysqlConfig = helper.Map{}

func InitDatabaseConfig(){
	MysqlConfig = helper.Map{
		"master" : helper.Map{
			"host" : helper.Env("db_host", "DB_MASTER"),
			"port" : helper.EnvInt("db_port", "DB_MASTER"),
			"username" : helper.Env("db_username", "DB_MASTER"),
			"password" : helper.Env("db_password", "DB_MASTER"),
			"dbname" : helper.Env("db_dbname", "DB_MASTER"),
		},
		"slave" : helper.Map{
			"host" : helper.Env("db_slave_host", "DB_SLAVE"),
			"port" : helper.EnvInt("db_slave_port", "DB_SLAVE"),
			"username" : helper.Env("db_slave_username", "DB_SLAVE"),
			"password" : helper.Env("db_slave_password", "DB_SLAVE"),
			"dbname" : helper.Env("db_slave_dbname", "DB_SLAVE"),
		},
	}
}

