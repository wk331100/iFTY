package request

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/wk331100/iFTY/system/helper"
	"strings"
)

func Input(ctx *fasthttp.RequestCtx) helper.Map {
	//Query
	query := string(ctx.URI().QueryString())
	params := helper.Map{}
	if query != "" {
		parseQuery(query, &params)
	}
	//Form Data
	if formData,err := ctx.MultipartForm(); err == nil{
		for key, item := range (*formData).Value{
			if len(item) == 1{
				params[key] = item[0]
			} else {
				params[key] = item
			}
		}
	}
	//form urlencoded
	formEncodeString := ctx.PostArgs().String()
	if formEncodeString != ""{
		parseQuery(formEncodeString, &params)
	}

	fmt.Println(params)
	return params
}

func parseQuery(query string, params *helper.Map) helper.Map {
	pair:=strings.Split(query, "&")
	if len(pair) > 1 {
		for _,val := range pair{
			b:=strings.Split(val, "=")
			(*params)[b[0]] = b[1]
		}
	} else {
		b:=strings.Split(query, "=")
		(*params)[b[0]] = b[1]
	}
	return *params
}


