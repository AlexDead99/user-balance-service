package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetOrders(t *testing.T) {
	var transferId int32 = 1
	orders, err := testQueries.ListOrders(context.Background(), transferId)
	require.NoError(t, err)
	require.NotEmpty(t, orders)
}

func TestDeleteOrder(t *testing.T) {
	var orderId int32 = 1
	err := testQueries.DeleteOrder(context.Background(), orderId)
	require.NoError(t, err)
}
