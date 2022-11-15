package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetService(t *testing.T) {
	availableService, err := testQueries.GetService(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, availableService)

	require.Equal(t, availableService.Name, "Buy")
}
