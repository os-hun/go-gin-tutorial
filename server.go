package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"go-gin-tutorial/database"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	database.DbInit()

	// Index
	r.GET("/", func(c *gin.Context) {
		todos := database.DbGetAll()
		c.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	// Create
	r.POST("/new", func(c *gin.Context) {
		text := c.PostForm("text")
		status := c.PostForm("status")
		database.DbInsert(text, status)
		c.Redirect(302, "/")
	})

	// Detail
	r.GET("/detail/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := database.DbGetOne(id)
		c.HTML(200, "detail.html", gin.H{
			"todo": todo,
		})
	})

	// Update
	r.POST("/update/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := c.PostForm("text")
		status := c.PostForm("status")
		database.DbUpdate(id, text, status)
		c.Redirect(302, "/")
	})

	// 削除確認
	r.GET("/delete/:id/check", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		todo := database.DbGetOne(id)
		c.HTML(200, "delete.html", gin.H{
			"todo": todo,
		})
	})

	// Delete
	r.POST("/delete/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		database.DbDelete(id)
		c.Redirect(302, "/")
	})

	// listen and serve on :8080
	r.Run()
}
