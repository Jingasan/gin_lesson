package main

// Ginのimport
import (
	"github.com/gin-gonic/gin"
)

func main() {
	// GinをReleaseモードに設定
	gin.SetMode(gin.ReleaseMode)
	// Engineインスタンスの作成
	engine := gin.Default()
	// POSTメソッド
	engine.POST("/hello", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
	// GETメソッド
	engine.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
	// PUTメソッド
	engine.PUT("/hello", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
	// DELETEメソッド
	engine.DELETE("/hello", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
	// ポート番号を指定
	engine.Run(":3000")
}
