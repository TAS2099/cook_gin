package main

import (
	"cook_gin/model"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// ルーターを準備
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	// ユーザー登録画面
	router.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "signup.html", gin.H{})
	})

	// ユーザー登録
	router.POST("/signup", func(c *gin.Context) {
		var users model.User
		// バリデーション
		if err := c.ShouldBind(&users); err != nil {
			c.HTML(400, "signup.html", gin.H{"err": err})
			c.Abort() // これ以下の処理をストップ
		} else {
			username := c.PostForm("username")
			password := c.PostForm("password")
			// 重複するユーザーを弾く
			if err := model.CheckUser(username, password); err != nil {
				c.HTML(400, "signup.html", gin.H{"err": err})
			}
			c.Redirect(302, "/")
		}
	})

	model.Init()

	// Index
	router.GET("/", func(ctx *gin.Context) {
		todos := model.SelectAll()
		ctx.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	// Create
	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		model.TodoInsert(text, status)
		ctx.Redirect(302, "/")
	})

	// Detail
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := model.SelectOne(id)
		ctx.HTML(200, "detail.html", gin.H{"todo": todo})
	})

	// Update
	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		model.Update(id, text, status)
		ctx.Redirect(302, "/")
	})

	// Delete
	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		model.Delete(id)
		ctx.Redirect(302, "/")
	})

	// サーバーを起動
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("サーバー起動に失敗", err)
	}
}
