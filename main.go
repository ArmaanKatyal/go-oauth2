package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ArmaanKatyal/go-oauth2/config"
	"github.com/ArmaanKatyal/go-oauth2/controllers"
	"github.com/ArmaanKatyal/go-oauth2/internal"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func main() {
	internal.InitializeRedis("localhost", "6379")

	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/health", health)

	config.GoogleConfig()
	e.GET("/google_login", controllers.GoogleLogin)
	e.GET("/google_callback", controllers.GoogleCallback)
	e.GET("/profile", controllers.Profile)

	sig_ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	<-sig_ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	e.Logger.Info("Server gracefully stopped")
}

func health(c echo.Context) error {
	req := c.Request()
	format := `
		<code>
		Protocol: %s<br>
		Host: %s<br>
		Remote Address: %s<br>
		Method: %s<br>
		Path: %s<br>
		</code>
	`
	return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path))
}
