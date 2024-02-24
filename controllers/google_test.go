package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ArmaanKatyal/go-oauth2/internal"
	"github.com/alicebob/miniredis/v2"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
}

func TestGenerateState(t *testing.T) {
	state, err := generateState()
	if assert.NoError(t, err) {
		assert.NotEmpty(t, state)
		assert.Equal(t, 64, len(state))
	}
}

func TestGenerateJWT(t *testing.T) {
	user := map[string]interface{}{
		"email":          "test",
		"name":           "test",
		"verified_email": true,
	}
	token, err := generateJWT(user)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, token)
	}
}

func TestGoogleLogin(t *testing.T) {
	s := miniredis.RunT(t)
	defer s.Close()

	internal.InitializeRedis(s.Host(), s.Port())

	req := httptest.NewRequest(http.MethodGet, "/google_login", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	if assert.NoError(t, GoogleLogin(ctx)) {
		assert.Equal(t, http.StatusFound, rec.Code)
		assert.NotEmpty(t, rec.Header().Get("Location"))
	}
}
