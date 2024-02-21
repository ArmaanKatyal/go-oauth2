package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ArmaanKatyal/go-oauth2/config"
	"github.com/ArmaanKatyal/go-oauth2/controllers"
	"github.com/ArmaanKatyal/go-oauth2/internal"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	internal.InitializeRedis("localhost", "6379")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/health", health)

	config.GoogleConfig()
	e.GET("/google_login", controllers.GoogleLogin)
	e.GET("/google_callback", controllers.GoogleCallback)
	e.GET("/profile", controllers.Profile)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func health(ctx echo.Context) error {
	return ctx.String(200, "OK")
}
