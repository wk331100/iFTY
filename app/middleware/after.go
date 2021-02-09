package middleware

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/wk331100/iFTY/system/log"
)

type After struct {}

func (a *After)Handle(ctx *fasthttp.RequestCtx) {
	fmt.Println("------after handle-----")
	log.Info("Request Log From After MiddleWare")
}
