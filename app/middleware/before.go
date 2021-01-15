package middleware

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type Before struct {}

func (b *Before)Handle(ctx *fasthttp.RequestCtx) {
	fmt.Println("------before handle-----")

}
