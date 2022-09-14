package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	//   c.JSON(http.StatusOK, gin.H{
	//     "message": "pong",
	//   })
	// })
	r.Any("/crud", func(ctx *gin.Context) {
		//获取请求方法和请求参数
		params := ctx.GetString("yes")
		fmt.Print("params-------------》: %s", params)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
