package main

import (
	"first/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", handlers.IndexHandler)
	router.GET("/:name", handlers.IndexHandler)
	router.Run()
}
