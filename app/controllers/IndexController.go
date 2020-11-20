package controllers

import (
	"github.com/valyala/fasthttp"
)

type Index struct {}

func (index *Index) Index(ctx *fasthttp.RequestCtx){
	ctx.Success("application/json",[]byte("Index Application/json"))
}

func (index *Index) Post(ctx *fasthttp.RequestCtx){
	ctx.Success("application/json",[]byte("Post Application/json"))
}