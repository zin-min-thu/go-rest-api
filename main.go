package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zin-min-thu/go-rest-api/db"
	"github.com/zin-min-thu/go-rest-api/models"
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

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080
}

func getHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Go REST API Project",
	})
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get events"})
		return
	}
	if len(events) == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "no events found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"events": events})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"event": event})
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}
