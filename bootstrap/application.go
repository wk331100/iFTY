package bootstrap

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/wk331100/iFTY/config"
	"github.com/wk331100/iFTY/route"
)

type Application struct {}

func (app *Application) bootstrap ()  {
	//加载系统配置文件
	config := config.ServerConfig
	//加载路由
	apiRoute := route.ApiRoute{}
	apiMap := apiRoute.Map()

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		fmt.Println("+++++++++")
		if len(apiMap) <= 0 {
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}

		requestMatch := false
		for _, item := range apiMap{
			if string(ctx.Path()) == item.Path && string(ctx.Method()) == item.Method{
				requestMatch = true
				item.Function.(func(*fasthttp.RequestCtx))(ctx)
			}
		}
		if !requestMatch {
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}

	}
	addr := fmt.Sprintf(":%d", config["Port"].(int))
	fasthttp.ListenAndServe(addr, requestHandler)
}

func (app *Application) Run () {
	app.bootstrap()
}

func fooHandler(ctx *fasthttp.RequestCtx){
	ctx.Success("application/json",[]byte("hello Application/json"))
}

func barHandler(ctx *fasthttp.RequestCtx){
	ctx.Success("text/html",[]byte("text/html"))
}

func errHandler(ctx *fasthttp.RequestCtx){
	ctx.Error("not found", fasthttp.StatusNotFound)
}