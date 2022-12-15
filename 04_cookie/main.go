package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// GETメソッド
	router.GET("/cookie", func(ctx *gin.Context) {
		// Cookieの取得
		cookie, err := ctx.Cookie("gin_cookie")
		// Cookieがない場合
		if err != nil {
			cookie = "NotSet"
			// Cookieの設定
			ctx.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.Run(":3000")
}
