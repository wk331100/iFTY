package controllers

import (
	"github.com/valyala/fasthttp"
	"github.com/wk331100/iFTY/system/helper"
	"github.com/wk331100/iFTY/system/http/request"
	"github.com/wk331100/iFTY/system/http/response"
)

type IndexController struct {}

func (index *IndexController) Index(ctx *fasthttp.RequestCtx){
	params := request.Input(ctx)
	name := params.String("name")
	age := params.Int("age")
	data := helper.Map{
		"Name" : name,
		"Age" : age,
	}
	response.Json(data, ctx)
}

func (index *IndexController) Post(ctx *fasthttp.RequestCtx){
	params := request.Input(ctx)
	name := params.String("name")
	age := params.Int("age")
	data := helper.Map{
		"Name" : name,
		"Age" : age,
	}
	response.Json(data, ctx)
}