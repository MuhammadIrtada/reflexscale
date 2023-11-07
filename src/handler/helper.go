package handler

import (
	"reflexscale/src/entitiy"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(statusCode, entity.HTTPResponse{
		Message:    message,
		IsSuccess:  true,
		Data:       data,
	})
}

func ErrorResponse(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, entity.HTTPResponse{
		Message:   err.Error(),
		IsSuccess: false,
		Data:      nil,
	})
}
