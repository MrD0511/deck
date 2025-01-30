package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

// StartServer starts the Gin server
func StartServer() {
	router := gin.Default()
	RegisterRoutes(router)

	log.Println("Server running on :8080")
	router.Run(":8081")
}
