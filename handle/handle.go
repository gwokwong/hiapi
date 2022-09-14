package handle

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//全局在这里分支
func handle(ctx *gin.Context) {
	fmt.Println("ctx---->")
	fmt.Println(ctx)
	method := ctx.Request.Method
	switch method {
	case "GET":
		fmt.Println("获取请求方法22222:", method)
		break

	}
	//获取请求方法
	fmt.Println("获取请求方法:", method)
	params := ""
	if method == "get" {
		params = ctx.GetString("yes")
	}
	//请求参数
	fmt.Println("params-------------》", params)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
