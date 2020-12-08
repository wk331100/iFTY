package controllers

import (
	"github.com/valyala/fasthttp"
	"github.com/wk331100/iFTY/app/services"
	"github.com/wk331100/iFTY/system/helper"
	"github.com/wk331100/iFTY/system/http/request"
	"github.com/wk331100/iFTY/system/http/response"
)

type IndexController struct {}

func (index *IndexController) List(ctx *fasthttp.RequestCtx){
	params := request.Input(ctx)
	name := params.String("name")
	age := params.Int("age")
	ids := params.String("ids")

	filter := helper.Map{}
	if name != "" {
		filter["name"] = name
	}
	if age != 0 {
		filter["age"] = age
	}
	if ids != "" {
		filter["ids"] = ids
	}

	result,errCode := new(services.TestService).List(filter)
	if errCode != nil{
		response.Error(errCode, ctx)
	} else {
		response.Json(result, ctx)
	}
}

func (index *IndexController) Info(ctx *fasthttp.RequestCtx){
	params := request.Input(ctx)
	id := params.String("id")

	filter := helper.Map{
		"id" : id,
	}
	result,errCode := new(services.TestService).Info(filter)
	if errCode != nil{
		response.Error(errCode, ctx)
	} else {
		response.Json(result, ctx)
	}
}

func (index *IndexController) Create(ctx *fasthttp.RequestCtx){
	params := request.Input(ctx)
	name := params.String("name")
	age := params.Int("age")
	data := helper.Map{
		"name" : name,
		"age" : age,
	}
	result,errCode := new(services.TestService).Create(data)
	if errCode != nil{
		response.Error(errCode, ctx)
	} else {
		response.Json(result, ctx)
	}
}

func (index *IndexController) Update(ctx *fasthttp.RequestCtx){
	params := request.Input(ctx)
	name := params.String("name")
	age := params.Int("age")
	id := params.Int("id")

	update := helper.Map{
		"name" : name,
		"age" : age,
	}
	filter := helper.Map{
		"id" : id,
	}
	result,errCode := new(services.TestService).Update(update, filter)
	if errCode != nil{
		response.Error(errCode, ctx)
	} else {
		response.Json(result, ctx)
	}
}


func (index *IndexController) Delete(ctx *fasthttp.RequestCtx){
	params := request.Input(ctx)
	id := params.String("id")

	filter := helper.Map{
		"id" : id,
	}
	result,errCode := new(services.TestService).Delete(filter)
	if errCode != nil{
		response.Error(errCode, ctx)
	} else {
		response.Json(result, ctx)
	}
}