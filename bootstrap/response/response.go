package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Json(param interface{}, c *gin.Context) {
	resp := map[string]interface{}{
		"Code" : 200,
		"Msg" : "成功",
		"Data" :  param,
	}
	c.JSON(http.StatusOK, resp)
}
