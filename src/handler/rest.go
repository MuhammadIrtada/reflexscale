package handler

import (
	"log"
	"os"
	"reflexscale/src/usecase"

	"github.com/gin-gonic/gin"
)

type rest struct {
	http *gin.Engine
	uc   *usecase.Usecase
}

func Init(uc *usecase.Usecase) *rest {
	r := &rest{}

	r.http = gin.Default()
	r.uc = uc

	r.RegisterMiddlewareAndRoutes()

	return r
}

func (r *rest) RegisterMiddlewareAndRoutes() {
	// Global middleware
	r.http.Use(r.CorsMiddleware())

	v1 := r.http.Group("api/v1")

	v1.POST("register", r.CreateUser)
	v1.POST("login", r.UserLogin)

	auth := v1.Group("", r.ValidateToken())
	{
		userGroup := auth.Group("user") 
		{
			userGroup.GET("", r.ReadAll)
			userGroup.GET(":id", r.ReadByID)
			userGroup.PATCH(":id", r.Update)
			userGroup.DELETE(":id", r.Delete)
		}
		
		hasilTestGroup := auth.Group("hasil-test")
		{
			hasilTestGroup.POST("", r.CreateHasilTest)
		}
	}

}

func (r *rest) Run() {
	port := ":8080"
	if os.Getenv("APP_PORT") != "" {
		port = ":" + os.Getenv("APP_PORT")
	}

	if err := r.http.Run(port); err != nil {
		log.Fatal(err)
	}
}
