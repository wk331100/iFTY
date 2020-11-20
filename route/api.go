package route

import (
	"github.com/wk331100/iFTY/app/controllers"
	"github.com/wk331100/iFTY/app/controllers/admin"
)

type ApiRoute struct {}

func (api *ApiRoute) Map() []Map{
	route := new(Route)

	//配置静态路由
	indexController := new(controllers.Index)
	route.Get("/hello", indexController.Index)
	route.Post("/hello", indexController.Post)

	//配置路由组
	route.Group("/admin", func() {
		userController := new(admin.UserController)
		route.Get("/getInfo", userController.GetInfo)
		route.Post("/create", userController.Create)
	})
	return route.Map
}




