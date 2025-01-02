package main

import (
	"myrestapi/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "timestamp": time.Now()})
	})

	server.GET("/events", getEvents)
	server.POST("/event", createEvent)

	server.Run(":8085")
}

func getEvents(c *gin.Context) {
	c.JSON(200, models.GetEvents())
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot parse Event JSON data"})
		return
	}

	event.Id = 1
	event.UserId = 1
	event.Save()
	c.JSON(http.StatusOK, gin.H{"message": "Event created", "event": event})
}
