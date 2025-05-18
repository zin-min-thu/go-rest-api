package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zin-min-thu/go-rest-api/models"
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
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080
}

func getHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Go REST API Project",
	})
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}
