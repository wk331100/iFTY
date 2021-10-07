package route

import (
	"blog/app/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Route struct {
	Method string
	Path string

	Func gin.HandlerFunc
}

var Routes []Route

func (r *Route) Map() []Route {
	userController := controller.UserController{}
	r.Register(http.MethodGet, "/", userController.Index)
	r.Register(http.MethodGet, "/info", userController.Info)
	r.Register(http.MethodGet, "/list", userController.List)
	return Routes
}

func (r *Route)Register(method, path string, target func(*gin.Context))  {
	route := Route{
		Method: method,
		Path:   path,
		Func:   target,
	}
	Routes = append(Routes, route)
}



