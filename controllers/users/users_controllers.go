package users

import (
	"fmt"
	"net/http"
	"strconv"

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
		restErr := errors.CreateNewBadReqError("Invalid JSON data")
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
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.CreateNewBadReqError("Invalid user id")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
