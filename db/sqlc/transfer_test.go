package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/PuneetBirdi/golang-bank/util"
)

func createRandomTransfer(t * testing.T, from_account Account, to_account Account) Transfer {
	arg := CreateTransferParams {
		FromAccountID: from_account.ID,
		ToAccountID: to_account.ID,
		Amount: util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)
	createRandomTransfer(t, from_account, to_account)
}

func TestGetTransfer(t *testing.T) {
	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)
	transfer := createRandomTransfer(t, from_account, to_account)

	arg := transfer.ID

	transferRes, err := testQueries.GetTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transferRes)

	require.Equal(t, transferRes.ID, transfer.ID)
	require.Equal(t, transferRes.FromAccountID, transfer.FromAccountID)
	require.Equal(t, transferRes.ToAccountID, transfer.ToAccountID)
	require.Equal(t, transferRes.Amount, transfer.Amount)
}

func TestListTransfers(t *testing.T) {
	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)
	
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, from_account, to_account)
	}

	args := ListTransfersParams {
		FromAccountID: from_account.ID,
		ToAccountID: to_account.ID,
		Limit: 5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.Equal(t, transfer.FromAccountID, from_account.ID)
		require.Equal(t, transfer.ToAccountID, to_account.ID)
		require.NotEmpty(t, transfer.Amount)
	}
}
