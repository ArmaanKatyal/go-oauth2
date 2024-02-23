package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeRedis(t *testing.T) {
	InitializeRedis("localhost", "6379")
	assert.NotNil(t, RedisClient.Conn)
}

func TestGetRedisClient(t *testing.T) {
	InitializeRedis("localhost", "6379")
	assert.NotNil(t, GetRedisClient())
}
