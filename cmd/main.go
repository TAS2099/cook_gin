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
		model.Insert(text, status)
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
