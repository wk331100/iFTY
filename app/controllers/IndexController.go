package controllers

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/wk331100/iFTY/app/models"
	"github.com/wk331100/iFTY/system/helper"
	"github.com/wk331100/iFTY/system/http/request"
	"github.com/wk331100/iFTY/system/http/response"
)

type IndexController struct {}

func (index *IndexController) Index(ctx *fasthttp.RequestCtx){
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


	result := new(models.TestModel).List(filter)
	fmt.Println(result)
	response.Json(result, ctx)
}

func (index *IndexController) Info(ctx *fasthttp.RequestCtx){
	params := request.Input(ctx)
	id := params.String("id")

	filter := helper.Map{
		"id" : id,
	}
	result := new(models.TestModel).Info(filter)
	response.Json(result, ctx)
}

func (index *IndexController) Post(ctx *fasthttp.RequestCtx){
	params := request.Input(ctx)
	name := params.String("name")
	age := params.Int("age")
	data := helper.Map{
		"name" : name,
		"age" : age,
	}
	fmt.Println(data)
	result := new(models.TestModel).Insert(data)
	response.Json(result, ctx)

}

func (index *IndexController) Put(ctx *fasthttp.RequestCtx){
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
	result := new(models.TestModel).Update(update, filter)
	response.Json(result, ctx)

}


func (index *IndexController) Delete(ctx *fasthttp.RequestCtx){
	params := request.Input(ctx)
	name := params.String("name")
	age := params.Int("age")

	filter := helper.Map{
		"name" : name,
		"age" : age,
	}
	result := new(models.TestModel).Delete(filter)
	response.Json(result, ctx)

}