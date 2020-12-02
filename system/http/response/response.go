package response

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/wk331100/iFTY/system/helper"
)

func Json(data helper.Map, ctx *fasthttp.RequestCtx)  {
	response := helper.Map{
		"Code" : fasthttp.StatusOK,
		"Msg" : "成功",
		"Data" : data,
	}
	jsonEncode, err := json.Marshal(response)
	if err != nil {
		panic("Json Encode Error")
	}
	ctx.Success("application/json", jsonEncode)
}