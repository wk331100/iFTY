package response

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/wk331100/iFTY/app/libs/errorCode"
	"github.com/wk331100/iFTY/system/helper"
)

func Json(data interface{}, ctx *fasthttp.RequestCtx)  {

	response := helper.Map{
		"Code" : errorCode.SUCCESS,
		"Msg" :  errorCode.GetErrorMessage(errorCode.SUCCESS),
		"Data" : data,
	}
	jsonEncode, err := json.Marshal(response)
	if err != nil {
		panic("Json Encode Error")
	}
	ctx.Success("application/json", jsonEncode)
}

func Error(code interface{}, ctx *fasthttp.RequestCtx)  {

	response := helper.Map{
		"Code" : code,
		"Msg" :  errorCode.GetErrorMessage(code.(int)),
		"Data" : "",
	}
	jsonEncode, err := json.Marshal(response)
	if err != nil {
		panic("Json Encode Error")
	}
	ctx.Success("application/json", jsonEncode)
}