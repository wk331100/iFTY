package helper

import (
	"github.com/Unknwon/goconfig"
	"github.com/wk331100/iFTY/system/global"
)

//从Env文件中读取String类型值
func Env(key, section string) string {
	configFile := ".env"
	if global.Env != "dev" {
		configFile += "_" +global.Env
	}
	env, err := goconfig.LoadConfigFile(configFile)
	if err == nil {
		value,err := env.GetValue(section, key)
		if err == nil{
			return value
		}
	}
	return ""
}

//从Env文件中读取Int类型的值
func EnvInt(key, section string) int{
	configFile := ".env"
	if global.Env != "dev" {
		configFile += "_" +global.Env
	}
	env, err := goconfig.LoadConfigFile(configFile)
	if err == nil {
		value,err := env.Int(section, key)
		if err == nil{
			return value
		}
	}
	return 0
}

