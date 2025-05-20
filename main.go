package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zin-min-thu/go-rest-api/db"
	"github.com/zin-min-thu/go-rest-api/routes"
)

func main() {

	db.InitDB()

	server := gin.Default()

	server.LoadHTMLGlob("templates/*")

	// server.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"title": "Go REST API Project",
	// 	})
	// })
	server.GET("/", getHome)

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}

func getHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Go REST API Project",
	})
}
