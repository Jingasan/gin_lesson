package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// リクエストボディのデータ型
type RequestBodyDataType struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// POSTメソッド処理関数
func postFunction(ctx *gin.Context) {
	var json RequestBodyDataType
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"firstname": json.Firstname, "lastname": json.Lastname})
}

func main() {
	router := gin.Default()
	// /hello
	// BodyData: {"firstname": "XXX", "lastname": "YYY"}
	router.POST("/hello", postFunction)
	router.Run(":3000")
}
