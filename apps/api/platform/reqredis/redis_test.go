package reqredis

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnect_Success(t *testing.T) {
	url := os.Getenv("REDIS_URL")
	if url == "" {
		t.Skip("REDIS_URL not set; skipping reqredis.Connect integration test")
	}

	client, err := Connect(context.Background(), url)
	require.NoError(t, err)
	defer client.Close()

	require.NoError(t, client.Ping(context.Background()).Err())
}

func TestConnect_InvalidURL(t *testing.T) {
	_, err := Connect(context.Background(), "not-a-valid-url")
	require.Error(t, err)
}

func TestConnect_Unreachable(t *testing.T) {
	_, err := Connect(context.Background(), "redis://127.0.0.1:1")
	require.Error(t, err)
}
