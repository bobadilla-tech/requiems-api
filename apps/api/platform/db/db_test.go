package db

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnect_Success(t *testing.T) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL not set; skipping db.Connect integration test")
	}

	pool, err := Connect(context.Background(), dsn)
	require.NoError(t, err)
	defer pool.Close()

	require.NoError(t, pool.Ping(context.Background()))
}

func TestConnect_InvalidDSN(t *testing.T) {
	_, err := Connect(context.Background(), "://not-a-valid-dsn")
	require.Error(t, err)
}

func TestConnect_UnreachableHost(t *testing.T) {
	_, err := Connect(context.Background(), "postgres://invalid-host-that-does-not-exist/db?sslmode=disable&connect_timeout=1")
	require.Error(t, err)
}
