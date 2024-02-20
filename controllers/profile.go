package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Claims struct {
	Email          string `json:"email"`
	Name           string `json:"name"`
	Verified_Email bool   `json:"verified_email"`
	jwt.RegisteredClaims
}

func Profile(ctx echo.Context) error {
	token := ExtractAuthToken(ctx)
	if token == "" {
		return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	claims := &Claims{}
	parsed, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !parsed.Valid {
		return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	message := fmt.Sprintf("<h1>Welcome %s</h1>", claims.Name)
	return ctx.HTML(http.StatusOK, message)
}

func ExtractAuthToken(ctx echo.Context) string {
	token, err := ctx.Request().Cookie("Authorization")
	if err != nil {
		return ""
	}
	return token.Value
}
