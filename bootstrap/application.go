package bootstrap

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/wk331100/iFTY/config"
	"github.com/wk331100/iFTY/route"
	"github.com/wk331100/iFTY/system/helper"
	"strconv"
)

type Application struct {}

func (app *Application) bootstrap ()  {
	//加载系统配置文件
	configs := config.ServerConfig
	//加载路由
	apiRoute := route.ApiRoute{}
	apiMap := apiRoute.Map()

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		fmt.Println("Request: ", string(ctx.Path()), " ", string(ctx.Method()))
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
	addr := fmt.Sprintf(":%d", configs["Port"].(int))
	app.PrintWelcome(configs["Port"].(int))
	fasthttp.ListenAndServe(addr, requestHandler)
}

func (app *Application) Run () {
	app.bootstrap()
}

func (app *Application) PrintWelcome (port int){
	fmt.Println("+-------------------------------------------------------+")
	fmt.Println("|                          iFTY                         |")
	fmt.Println("|       Is A Web Api Framework Short For Infinity       |")
	fmt.Println("|-------------------------------------------------------|")
	fmt.Println("| Status: ",helper.Green("Running!"),"                                    |")
	fmt.Println("| Listening Port: ", helper.Green(strconv.Itoa(port)), "                                |")
	fmt.Println("+-------------------------------------------------------+")
}