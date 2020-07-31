package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ping function
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
