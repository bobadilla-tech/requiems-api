package mx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestService_LookupBatch(t *testing.T) {
	t.Parallel()
	svc := NewService()

	domains := []string{
		"yahoo.com", "not-exist-ailxz.com", "icloud.com",
	}

	results := svc.LookupBatch(context.Background(), domains)

	require.Equal(t, 3, len(results))
	require.True(t, results[0].Found)
	require.False(t, results[1].Found)
	require.True(t, results[2].Found)
}

func TestService_LookupBatch_ContextCancelled(t *testing.T) {
	t.Parallel()
	svc := NewService()

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel immediately

	results := svc.LookupBatch(ctx, []string{"gmail.com", "outlook.com"})

	// all should fail because context is cancelled
	for _, item := range results {
		require.False(t, item.Found)
		require.NotEmpty(t, item.Error)
	}
}
