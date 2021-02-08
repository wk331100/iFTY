package admin

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/wk331100/iFTY/app/models"
	"github.com/wk331100/iFTY/system/helper"
	"github.com/wk331100/iFTY/system/http/response"
)

type UserController struct {}

func (user *UserController) List(ctx *fasthttp.RequestCtx)  {
	fmt.Println("++++++++++++++")
	result,_ :=  new(models.UserModel).Instance().List(helper.Map{})
	response.Json(result, ctx)
}

func (user *UserController) Test(ctx *fasthttp.RequestCtx)  {
	result,_ :=  new(models.UserModel).Instance().GetTest()
	response.Json(result, ctx)
}

func (user *UserController) Create(ctx *fasthttp.RequestCtx)  {
	fmt.Println("hello Admin User Create")
}

