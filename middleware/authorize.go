package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type SignedDetails struct {
	Username string
	ID       int
	jwt.StandardClaims
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Authorize(c *gin.Context) {

	signedToken := c.GetHeader("Authorization")
	if signedToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})

		return
	}

	if strings.Split(signedToken, " ")[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})

		return
	}

	splitToken := strings.Split(signedToken, "Bearer ")

	token, err := jwt.ParseWithClaims(
		splitToken[1],
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})

		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Invalid token",
		})

		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Token is expired",
		})

		return
	}

	c.Set("UserID", claims.ID)

	c.Next()
}
