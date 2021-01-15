package config

import (
	Middleware "github.com/wk331100/iFTY/app/middleware"
	"github.com/wk331100/iFTY/system/helper"
	Mid "github.com/wk331100/iFTY/system/middleware"
)

type MiddlewareContainer struct {}

func (this *MiddlewareContainer) Map() helper.Map {
	return helper.Map{
		"before" : this.BeforeMap(),
		"after" : this.AfterMap(),
	}
}

//前置中间件
func (this *MiddlewareContainer) BeforeMap() []Mid.Map {
	middleware := new(Mid.Middleware)
	middleware.Register("before", new(Middleware.Before).Handle)
	return middleware.Map
}

//后置中间件
func (this *MiddlewareContainer) AfterMap() []Mid.Map {
	middleware := new(Mid.Middleware)
	middleware.Register("after", new(Middleware.After).Handle)
	return middleware.Map
}
