package handler

import (
	"net/http"
	entity "reflexscale/src/entitiy"

	"github.com/gin-gonic/gin"
)

func (r *rest) CreateUser(c *gin.Context) {

	var userInput entity.UserRegisterInputParam

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	createUser, err := r.uc.User.Register(&userInput)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	SuccessResponse(c, 201, "Create user success", createUser)
}