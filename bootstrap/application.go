package bootstrap

import (
	"flag"
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/wk331100/iFTY/config"
	"github.com/wk331100/iFTY/route"
	"github.com/wk331100/iFTY/system/global"
	"github.com/wk331100/iFTY/system/helper"
	Mid "github.com/wk331100/iFTY/system/middleware"
	"strconv"
)

type Application struct {}

func (app *Application) bootstrap ()  {
	//设置系统环境
	flag.StringVar(&global.Env, "e", "dev", "Running Environment")
	flag.Parse()

	//初始化环境配置
	config.InitConfig()

	//加载系统配置文件
	configs := config.ServerConfig
	//加载路由
	apiMap := new(route.ApiRoute).Map()
	middlewareMap := new(config.MiddlewareContainer).Map()

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		fmt.Println("Request: ", string(ctx.Path()), " ", string(ctx.Method()))
		if len(apiMap) <= 0 {
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}

		//执行前置中间件
		beforeMiddlewareMap := middlewareMap["before"].([]Mid.Map)
		for _, item := range beforeMiddlewareMap{
			item.Function.(func(*fasthttp.RequestCtx))(ctx)
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

		//执行后置中间件
		afterMiddlewareMap := middlewareMap["after"].([]Mid.Map)
		for _, item := range afterMiddlewareMap{
			item.Function.(func(*fasthttp.RequestCtx))(ctx)
		}

	}
	addr := fmt.Sprintf(":%d", configs["Port"].(int))
	showParams := helper.Map{
		"port" : configs["Port"].(int),
	}
	app.PrintWelcome(showParams)
	fasthttp.ListenAndServe(addr, requestHandler)
}

func (app *Application) Run () {
	app.bootstrap()
}

func (app *Application) PrintWelcome (params helper.Map){
	fmt.Println("+-------------------------------------------------------+")
	fmt.Println("|                          iFTY                         |")
	fmt.Println("|       Is A Web Api Framework Short For Infinity       |")
	fmt.Println("|-------------------------------------------------------|")
	fmt.Println("| Status: ",helper.Green("Running!"),"                                    |")
	fmt.Println("| Listening Port: ", helper.Green(strconv.Itoa(params["port"].(int))), "                                |")
	fmt.Println("| Running Environment: ", helper.Green(global.Env), "                           |")
	fmt.Println("+-------------------------------------------------------+")
}