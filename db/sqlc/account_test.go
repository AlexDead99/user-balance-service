package db

import (
	"context"
	"testing"

	"github.com/AlexDead99/user-balance-service/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Accounts {
	arg := CreateAccountParams{
		Owner:   utils.CreateOwner(),
		Balance: utils.CreateBalance(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)

	require.NotZero(t, account.AccountID)
	require.NotZero(t, account.CreatedAt)

	return account
}
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	createdUser := createRandomAccount(t)

	availableUser, err := testQueries.GetAccount(context.Background(), createdUser.AccountID)
	require.NoError(t, err)
	require.NotEmpty(t, availableUser)

	require.Equal(t, createdUser.AccountID, availableUser.AccountID)
	require.Equal(t, createdUser.Balance, availableUser.Balance)
	require.Equal(t, createdUser.Owner, availableUser.Owner)
}

func TestUpdateAccount(t *testing.T) {
	createdUser := createRandomAccount(t)

	newBalance := createdUser.Balance + 15.0
	updatedParams := UpdateAccountParams{
		AccountID: createdUser.AccountID,
		Balance:   newBalance,
	}

	updatedUser, err := testQueries.UpdateAccount(context.Background(), updatedParams)

	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)

	require.Equal(t, createdUser.AccountID, updatedUser.AccountID)
	require.Equal(t, newBalance, updatedUser.Balance)
	require.Equal(t, createdUser.Owner, updatedUser.Owner)
}
