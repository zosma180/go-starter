package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zosma180/go-poc/dto"
	"github.com/zosma180/go-poc/service"
)

func CreateUser(c *gin.Context) {

	var model dto.CreateUserRequest

	if bind := c.ShouldBind(&model); bind != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("Invalid json provided: %s", bind))
	} else if err := service.CreateUser(model); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	} else {
		c.JSON(http.StatusCreated, model)
	}

}

func GetUserList(c *gin.Context) {
	list := service.GetUserList()
	c.JSON(http.StatusOK, list)
}
