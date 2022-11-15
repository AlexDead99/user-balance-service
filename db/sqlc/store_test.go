package db

import (
	"context"
	"testing"

	"github.com/AlexDead99/user-balance-service/utils"
	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	n := 5
	errs := make(chan error)
	results := make(chan TransferTxResult)
	for i := 0; i < n; i++ {
		txParams := TransferTxParams{
			UserId:      int32(i),
			Products:    []*ProductsParams{{3, 1}},
			ServiceId:   1,
			Description: utils.CreateDescription(),
		}
		go func() {
			result, err := store.TransferTx(context.Background(), txParams)

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		require.Equal(t, result.Success, true)
	}

}

func TestTransferTxWithInvalidProductId(t *testing.T) {
	store := NewStore(testDB)

	txParams := TransferTxParams{
		UserId:      int32(1),
		Products:    []*ProductsParams{{100, 1}},
		ServiceId:   1,
		Description: utils.CreateDescription(),
	}

	_, err := store.TransferTx(context.Background(), txParams)
	require.NotEmpty(t, err)
}
