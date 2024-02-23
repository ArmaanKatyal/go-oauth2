package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func TestExtractAuthToken_MissingToken(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	want := ""
	assert.Equal(t, want, ExtractAuthToken(ctx))
}

func TestExtractAuthToken_GivenToken(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.AddCookie(&http.Cookie{
		Name:  "Authorization",
		Value: "some-value",
	})
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	want := "some-value"
	assert.Equal(t, want, ExtractAuthToken(ctx))
}

func TestProfile_MissingToken(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	if assert.NoError(t, Profile(ctx)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Equal(t, "\"Unauthorized\"\n", rec.Body.String())
	}
}

func TestProfile_WrongToken(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.AddCookie(&http.Cookie{
		Name:  "Authorization",
		Value: "some-value",
	})
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	if assert.NoError(t, Profile(ctx)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Equal(t, "\"Unauthorized\"\n", rec.Body.String())
	}
}

func TestProfile_ExpiredToken(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	key := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":          "test",
		"name":           "test",
		"verified_email": false,
		"exp":            time.Now().Add(-24 * time.Hour).Unix(),
	})
	signed_token, err := token.SignedString([]byte(key))
	if err != nil {
		t.Fatal(err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "Authorization",
		Value: signed_token,
	})
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	if assert.NoError(t, Profile(ctx)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Equal(t, "\"Unauthorized\"\n", rec.Body.String())
	}
}

func TestProfile_ValidToken(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	key := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":          "test@test.com",
		"name":           "test",
		"verified_email": false,
		"exp":            time.Now().Add(24 * time.Hour).Unix(),
	})
	signed_token, err := token.SignedString([]byte(key))
	if err != nil {
		t.Fatal(err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "Authorization",
		Value: signed_token,
	})
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	if assert.NoError(t, Profile(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "test@test.com")
	}
}
