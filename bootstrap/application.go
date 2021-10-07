package bootstrap

import (
	"blog/config"
	"blog/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Application struct {
}

func (app *Application)Run()  {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
 	routes := route.Route{}
 	routeMap := routes.Map()
 	for _, routeItem := range routeMap {
		switch routeItem.Method {
		case http.MethodGet:
			fmt.Println(routeItem)
			r.GET(routeItem.Path, routeItem.Func)
		case http.MethodPost:
			r.POST(routeItem.Path, routeItem.Func)
		}
	 }

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	config.InitServerConfig()
	serverConfig := config.ServerConfig
	servStr := fmt.Sprintf(":%d", serverConfig["port"])
	r.Run(servStr)
}