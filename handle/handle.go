package handle

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//全局在这里分支
func Handle(ctx *gin.Context) {
	//获取请求方法后具体去执行方法
	fmt.Println("ctx---->")
	fmt.Println(ctx)
	method := ctx.Request.Method
	switch method {
	case "GET":
		fmt.Println("获取请求方法:", method)
		handleByGet(ctx)
		break
	case "POST":
		fmt.Println("获取请求方法:", method)
		handleByPost(ctx)
		break
	default:
		break
	}

	//获取请求方法
	// fmt.Println("获取请求方法:", method)
	// params := ""
	// if method == "get" {
	// 	params = ctx.GetString("yes")
	// }
	// //请求参数
	// fmt.Println("params-------------》", params)

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "wei",
	// })
}

//get方法的参数执行逻辑
func handleByGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"menthod": "GET",
	})
}

//post方法的参数执行逻辑
func handleByPost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"menthod": "POST",
	})
}
