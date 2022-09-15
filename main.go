package main

import (
	"example.com/m/v2/db"
	"example.com/m/v2/handle"
	"github.com/gin-gonic/gin"
)

func main() {

	db.Init("hiapi", "hiapi:5c7AKd72tCzZ4HiN@tcp(82.157.62.190:3306)")

	r := gin.Default()
	r.POST("/get", func(ctx *gin.Context) {
		handle.GetHandler(ctx)
	})

	r.POST("/post", func(ctx *gin.Context) {
		handle.PostHandler(ctx)
	})

	r.POST("/delete", func(ctx *gin.Context) {
		handle.DeleteHandler(ctx)
	})

	r.POST("/put", func(ctx *gin.Context) {
		handle.PutHandler(ctx)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
