package config

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2/google"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

}

func TestGoogleConfig(t *testing.T) {
	assert.Contains(t, GoogleConfig().RedirectURL, "google_callback")
	assert.Equal(t, os.Getenv("CLIENT_ID"), GoogleConfig().ClientID)
	assert.Equal(t, os.Getenv("CLIENT_SECRET"), GoogleConfig().ClientSecret)
	assert.Equal(t, len(GoogleConfig().Scopes), 2)
	assert.Equal(t, google.Endpoint, GoogleConfig().Endpoint)
}
