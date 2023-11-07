package handler

import (
	"net/http"
	entity "reflexscale/src/entitiy"

	"github.com/gin-gonic/gin"
)

func (r *rest) CreateHasilTest(c *gin.Context) {
	userId := c.MustGet("id").(float64)

	var hasilTestInput entity.HasilTestInputParam

	err := c.ShouldBindJSON(&hasilTestInput)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	createHasilTest, err := r.uc.HasilTest.CreateHasilTest(uint(userId), &hasilTestInput)
	if err != nil {
		ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	SuccessResponse(c, http.StatusCreated, "Create user success", createHasilTest)
}