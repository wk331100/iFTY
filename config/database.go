package config

import "github.com/wk331100/iFTY/system/helper"

var MysqlConfig = helper.Map{
	"master" : helper.Map{
		"host" : "192.168.126.100",
		"port" : 3306,
		"username" : "root",
		"password" : "Wk331100!",
		"dbname" : "blog",
	},
	"slave" : helper.Map{
		"host" : "192.168.126.100",
		"port" : 3306,
		"username" : "root",
		"password" : "Wk331100!",
		"dbname" : "blog",
	},
}
