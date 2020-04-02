package main

import "github.com/gin-gonic/gin"

func sayHi(ctx *gin.Context) {
	ctx.JSON(200,gin.H{
		"name":"james",
		"age":23,
	})
}
/**
	通过gin框架能够快速开启一个http server
 */
func main() {
	r := gin.Default()
	r.GET("/hi",sayHi)
	r.Run(":9090")
}
