package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// リクエストボディのデータ型
type RequestBodyDataType struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()
	// ミドルウェア登録
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	// ログインAPI
	r.POST("/login", func(c *gin.Context) {
		fmt.Println("> Login")
		// リクエストボディの取得
		var json RequestBodyDataType
		if err := c.ShouldBindJSON(&json); err != nil {
			// リクエストボディが規定の型を満たさない場合のエラー処理
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// セッションの取得
		session := sessions.Default(c)
		fmt.Println(session.Get("loginUser"))
		if session.Get("loginUser") == json.Username {
			c.String(http.StatusOK, "ログイン済")
			return
		}
		// セッションの作成
		session.Set("loginUser", json.Username)
		session.Save()
		c.String(http.StatusOK, "ログイン完了")
	})
	// ログアウトAPI
	r.GET("/logout", func(c *gin.Context) {
		fmt.Println("> Logout")
		// セッション取得
		session := sessions.Default(c)
		fmt.Println(session.Get("loginUser"))
		if session.Get("loginUser") == nil {
			c.String(http.StatusOK, "未ログイン")
			return
		}
		// セッションの破棄
		session.Clear()
		session.Save()
		c.String(http.StatusOK, "ログアウト完了")
	})
	r.Run(":3000") // listen and server on 0.0.0.0:3000
}
