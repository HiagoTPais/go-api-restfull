package controller

import (
	"fmt"
	"go-api-restfull/src/configuration/validation"
	"go-api-restfull/src/controller/model/request"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		fmt.Println("There are some incorrect filds, error=%s\n", err.Error())

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)

		return
	}

	fmt.Println(userRequest)
}
