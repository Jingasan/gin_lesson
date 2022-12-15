package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	// templatesフォルダ下のHTMLファイルを取得
	router.LoadHTMLGlob("templates/*")
	// GETメソッド
	router.GET("/index", func(c *gin.Context) {
		// HTMLにtitleを埋め込み、レスポンスとして返す
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":3000")
}
