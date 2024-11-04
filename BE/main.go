package main

import (
	"github.com/gin-contrib/cors"
	"log"
	"wan-api-kol-event/Initializers"
	"wan-api-kol-event/registry"

	"github.com/gin-gonic/gin"
)

func init() {
	Initializers.LoadEnvironmentVariables()
	Initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Config CORS
	r.Use(cors.Default())

	// Define routes
	registry.RegisterRoutes(r)

	// Run Gin server
	if err := r.Run(":8081"); err != nil {
		log.Println("Failed to start server")
	}
}
