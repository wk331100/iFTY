package controller

import (
	"iFTY/bootstrap/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserController struct {

}

func (c *UserController) Index(ctx *gin.Context){
	fmt.Println("111111111111")
	response.Json("Hello User", ctx)
}


func (c *UserController) List(ctx *gin.Context){

	response.Json("list", ctx)
}

func (c *UserController) Info(ctx *gin.Context) {
	fmt.Println("2222222222")
	data := map[string]interface{}{
		"uid" : 10010,
		"username" : "zhangsan",
		"age" : 18,
	}
	response.Json(data, ctx)
}

