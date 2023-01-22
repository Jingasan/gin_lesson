package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// GinをReleaseモードに設定
	gin.SetMode(gin.ReleaseMode)
	// Engineインスタンスの作成
	engine := gin.Default()
	// /user/john にはマッチするが、/user/ や /user にはマッチしない
	engine.GET("/user/:name", func(c *gin.Context) {
		// URLパラメータの取得
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	// /user/john/ と /user/john/send にマッチする
	engine.GET("/user/:name/*action", func(c *gin.Context) {
		// URLパラメータの取得
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	engine.Run(":3000")
}
