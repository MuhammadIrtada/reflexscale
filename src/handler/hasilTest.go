package handler

import (
	"net/http"
	"reflexscale/sdk/custome_time"
	"reflexscale/src/entity"

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

	SuccessResponse(c, http.StatusCreated, "Create user success", convertHasilTestResponse(*createHasilTest))
}

func convertHasilTestResponse(hasilTest entity.HasilTest) entity.HasilTestResponse {
	return entity.HasilTestResponse{
		ID:        hasilTest.ID,
		Poin:      hasilTest.Poin,
		CreatedAt: custome_time.ConvertTimeFormat(hasilTest.CreatedAt.String()),
	}
}