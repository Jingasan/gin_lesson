package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware1
func MyCustomMiddleWare1() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 本処理前に実行
		fmt.Println("> MyCustomMiddleWare1 Start")
		// アクセス元のIPアドレス取得
		fmt.Println("clientIP: ", context.ClientIP())
		// 本処理に入る
		context.Next()
		// 本処理後に実行
		fmt.Println("> MyCustomMiddleWare1 End")
	}
}

// Middleware2
func MyCustomMiddleWare2() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 本処理前に実行
		fmt.Println("> MyCustomMiddleWare1 Start")
		// 値を本処理に渡す
		context.Set("value", 123)
		// 本処理に入る
		context.Next()
		// 本処理後に実行
		fmt.Println("> MyCustomMiddleWare1 End")
	}
}

func main() {
	engine := gin.Default()
	// Middlewareの登録
	engine.Use(MyCustomMiddleWare1())
	engine.Use(MyCustomMiddleWare2())
	// 本処理
	engine.GET("/ping", func(context *gin.Context) {
		fmt.Println("> Route handler")
		// Middlewareから値を取得
		fmt.Println(context.Get("value"))
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
