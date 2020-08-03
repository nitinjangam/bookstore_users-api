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
	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		//TODO: Handle error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("x-public") == "true"))
}

//GetUser function
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.CreateNewBadReqError("Invalid user id")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.UserService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	//c.JSON(http.StatusOK, user)
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("x-public") == "true"))
}

//UpdateUser function
func UpdateUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.CreateNewBadReqError("Invalid user id")
		c.JSON(err.Status, err)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: Handle error
		restErr := errors.CreateNewBadReqError("Invalid JSON data")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	result, saveErr := services.UserService.UpdateUser(isPartial, user)
	if saveErr != nil {
		//TODO: Handle error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("x-public") == "true"))
}

//DeleteUser function
func DeleteUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.CreateNewBadReqError("Invalid user id")
		c.JSON(err.Status, err)
		return
	}
	if err := services.UserService.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

//SearchUser function
func SearchUser(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UserService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("x-public") == "true"))
}
