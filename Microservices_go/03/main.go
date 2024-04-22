package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin engine
	r := gin.Default()

	// Define a route and its handler
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// Run the server
	r.Run(":8080")
}
