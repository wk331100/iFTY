package middleware

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type After struct {}

func (a *After)Handle(ctx *fasthttp.RequestCtx) {
	fmt.Println("------after handle-----")

}
