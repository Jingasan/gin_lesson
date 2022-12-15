package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/hello", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
	router.PUT("/hello", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
	router.DELETE("/hello", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
	router.Run(":3000")
}
