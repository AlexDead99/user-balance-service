package db

import (
	"context"
	"testing"
)

func TestCreateAccount(t *testing.T){
	arg:=CreateAccountParams{
		Owner: "tom",
		Balance: 100
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)

	require.NotZero(t, account.AccountID)
	require.NotZero(t, account.CreatedAt)

}

