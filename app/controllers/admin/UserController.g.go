package admin

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/wk331100/iFTY/app/libs/errorCode"
	"github.com/wk331100/iFTY/app/models"
	errors "github.com/wk331100/iFTY/system/error"
	"github.com/wk331100/iFTY/system/helper"
	"github.com/wk331100/iFTY/system/http/request"
	"github.com/wk331100/iFTY/system/http/response"
	"github.com/wk331100/iFTY/system/validator"
)

type UserController struct {}

func (user *UserController) List(ctx *fasthttp.RequestCtx)  {
	params := request.Input(ctx)
	validate,errCode := validator.Make(params, helper.Map{
		"name" 	: "required|between:4,12",
		"age" 	: "required|int",
	})

	if errCode != errors.EMPTY || validate.Fails(){
		response.Error(errorCode.ERR_PARAMS, ctx)
		return
	}

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

