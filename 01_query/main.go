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
	router := gin.Default()
	// /hello?firstname=XXX&lastname=YYY
	router.GET("/hello", getFunction)
	router.Run(":3000")
}
