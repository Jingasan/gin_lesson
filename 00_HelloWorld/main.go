package main

// Ginのimport
import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// POSTメソッド
	router.POST("/hello", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
	// GETメソッド
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
	// PUTメソッド
	router.PUT("/hello", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
	// DELETEメソッド
	router.DELETE("/hello", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
	// ポート番号を指定
	router.Run(":3000")
}
