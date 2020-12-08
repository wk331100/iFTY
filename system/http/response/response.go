package response

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/wk331100/iFTY/app/libs/ErrorCode"
	"github.com/wk331100/iFTY/system/helper"
)

func Json(data interface{}, ctx *fasthttp.RequestCtx)  {

	response := helper.Map{
		"Code" : ErrorCode.SUCCESS,
		"Msg" :  ErrorCode.GetErrorMessage(ErrorCode.SUCCESS),
		"Data" : data,
	}
	jsonEncode, err := json.Marshal(response)
	if err != nil {
		panic("Json Encode Error")
	}
	ctx.Success("application/json", jsonEncode)
}

func Error(errorCode interface{}, ctx *fasthttp.RequestCtx)  {

	response := helper.Map{
		"Code" : errorCode,
		"Msg" :  ErrorCode.GetErrorMessage(errorCode.(int)),
		"Data" : "",
	}
	jsonEncode, err := json.Marshal(response)
	if err != nil {
		panic("Json Encode Error")
	}
	ctx.Success("application/json", jsonEncode)
}