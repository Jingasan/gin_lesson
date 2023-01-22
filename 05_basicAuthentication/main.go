package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// レスポンスとなる機密データ
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	// Engineインスタンスの作成
	engine := gin.Default()

	// gin.BasicAuth() ミドルウェアを使用したグループ
	// gin.Accounts は map[string]string へのショートカットである
	authorized := engine.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"john":   "4321",
	}))

	// 上記ミドルウェアにより、エンドポイントは [ip]:[port]/admin/secrets となる
	authorized.GET("/secrets", func(ctx *gin.Context) {
		// BasicAuth ミドルウェアで設定されたユーザー名にアクセスする
		user := ctx.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			ctx.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
	engine.Run(":3000")
}
