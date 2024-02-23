package controllers

import (
	"testing"

	"github.com/joho/godotenv"
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

}
