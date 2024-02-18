package controllers

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/ArmaanKatyal/go-oauth/config"
	"github.com/ArmaanKatyal/go-oauth/internal"
	"github.com/labstack/echo/v4"
)

func generateState() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(b)
	return hex.EncodeToString(hash[:]), nil
}

func GoogleLogin(ctx echo.Context) error {
	state, err := generateState()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to generate state")
	}

	err = internal.GetRedisClient().Set(context.Background(), state, "true", 10*time.Minute).Err()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to store state")
	}

	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL(state)
	ctx.Redirect(http.StatusFound, url)
	return nil
}

func GoogleCallback(ctx echo.Context) error {
	receivedState := ctx.QueryParam("state")

	stateExists, err := internal.GetRedisClient().Get(context.Background(), receivedState).Result()
	if err != nil || stateExists != "true" {
		return ctx.JSON(http.StatusInternalServerError, "Invalid or expired state")
	}

	internal.GetRedisClient().Del(context.Background(), receivedState)

	code := ctx.QueryParam("code")
	googlecon := config.GoogleConfig()
	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Code-Token exchange failed")
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get user info")
	}
	defer resp.Body.Close()

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to read user info")
	}

	// convert the response into a json object
	var user map[string]interface{}
	err = json.Unmarshal(userData, &user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to unmarshal user info")
	}

	return ctx.JSON(http.StatusOK, user)
}
