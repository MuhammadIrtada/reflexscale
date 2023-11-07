package handler

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (r *rest) CorsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})
}

func (r *rest) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		if bearerToken == "" {
			ErrorResponse(c, http.StatusUnauthorized, errors.New("token not found"))
			c.Abort()
			return
		}

		bearerToken = bearerToken[7:] // menghilangkan Bearer
		tokenExtract, err := jwt.Parse(bearerToken, tokenExtract)
		if err != nil {
			ErrorResponse(c, http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		if claims, ok := tokenExtract.Claims.(jwt.MapClaims); ok && tokenExtract.Valid {
			userId := claims["id"]

			c.Set("id", userId)
			c.Next()

			return
		}
		ErrorResponse(c, http.StatusForbidden, errors.New("invalid token"))
		c.Abort()
	}
}

func tokenExtract(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, jwt.ErrSignatureInvalid
	}

	return []byte(os.Getenv("SECRET_KEY")), nil
}
