package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.LoadHTMLGlob("templates/*")

	// server.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"title": "Go REST API Project",
	// 	})
	// })
	server.GET("/", getHome)

	server.GET("/events", getEvents)

	server.Run(":8080") // localhost:8080
}

func getHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Go REST API Project",
	})
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Events"})
}
