package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nitinjangam/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

//StartApplication function
func StartApplication() {
	mapUrls()
	logger.Info("About to start the application...")
	router.Run(":8080")
}
