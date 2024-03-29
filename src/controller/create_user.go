package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joaovitorcf97/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/joaovitorcf97/meu-primeiro-crud-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrenct filds, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
