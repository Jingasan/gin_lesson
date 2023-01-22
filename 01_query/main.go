package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GETメソッド処理関数
func getFunction(ctx *gin.Context) {
	// クエリの取得
	firstname := ctx.DefaultQuery("firstname", "Guest")
	lastname := ctx.Query("lastname")
	ctx.JSON(http.StatusOK, gin.H{
		"firstname": firstname,
		"lastname":  lastname,
	})
}

func main() {
	// GinをReleaseモードに設定
	gin.SetMode(gin.ReleaseMode)
	// Engineインスタンスの作成
	engine := gin.Default()
	// /hello?firstname=XXX&lastname=YYY
	engine.GET("/hello", getFunction)
	engine.Run(":3000")
}
