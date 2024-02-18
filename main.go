package main

import (
	"github.com/ArmaanKatyal/go-oauth2/config"
	"github.com/ArmaanKatyal/go-oauth2/controllers"
	"github.com/ArmaanKatyal/go-oauth2/internal"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	internal.InitializeRedis("redis", "6379")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/health", health)

	config.GoogleConfig()
	e.GET("/google_login", controllers.GoogleLogin)
	e.GET("/google_callback", controllers.GoogleCallback)

	e.Logger.Fatal(e.Start(":8080"))
}

func health(ctx echo.Context) error {
	return ctx.String(200, "OK")
}
