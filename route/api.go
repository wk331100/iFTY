package route

import (
	"github.com/wk331100/iFTY/app/controllers"
	"github.com/wk331100/iFTY/app/controllers/admin"
	Route "github.com/wk331100/iFTY/system/route"
)

type ApiRoute struct {}

func (api *ApiRoute) Map() []Route.Map {
	route := new(Route.Route)

	//配置静态路由
	indexController := new(controllers.IndexController)
	route.Get("/test", indexController.Index)
	route.Get("/info", indexController.Info)
	route.Post("/test", indexController.Post)
	route.Put("/test", indexController.Put)
	route.Delete("/test", indexController.Delete)

	//配置路由组
	route.Group("/admin", func() {
		userController := new(admin.UserController)
		route.Get("/getInfo", userController.GetInfo)
		route.Post("/create", userController.Create)
	})
	return route.Map
}




