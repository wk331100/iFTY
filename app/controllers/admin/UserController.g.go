package admin

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type UserController struct {}

func (user *UserController) GetInfo(ctx *fasthttp.RequestCtx)  {
	fmt.Println("hello Admin User GetInfo")
}

func (user *UserController) Create(ctx *fasthttp.RequestCtx)  {
	fmt.Println("hello Admin User Create")
}

