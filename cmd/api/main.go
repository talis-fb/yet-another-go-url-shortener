package main

import (
	"github.com/gin-gonic/gin"
	"github.com/talis-fb/yet-another-go-url-shortener/internal"
)

func main() {
	router := gin.Default()

	internal.SetupRoutes(router)

	router.Run(":8085")
}
