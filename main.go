package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hebertbanes/get-status/handlers"
)

func main() {
	// router instance
	r := gin.Default()

	// ROUTES
	// status
	r.POST("/status", handlers.GetStatusHandler)

	// run
	r.Run()
}
