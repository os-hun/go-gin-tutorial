package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-tutorial/database"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	
	database.DbInit()

	data := "Hello Go/Gin!!"
	r.GET("/ping", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"data": data,
		})
	})

	// listen and serve on :8080
	r.Run()
}
