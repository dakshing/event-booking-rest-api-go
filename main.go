package main

import (
	"github.com/dakshing/event-booking-rest-api-go/db"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db.InitDB()

	server := gin.Default()

	server.Run(":8080") // Start the server on port 8080

}
