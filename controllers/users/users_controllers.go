package users

import (
	"fmt"
	"net/http"

	"github.com/nitinjangam/bookstore_users-api/services"

	"github.com/nitinjangam/bookstore_users-api/domain/users"

	"github.com/gin-gonic/gin"
	"github.com/nitinjangam/bookstore_users-api/utils/errors"
)

//CreateUser function
func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: Handle error
		restErr := errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: Handle error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

//GetUser function
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Me!")
}
