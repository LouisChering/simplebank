package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	ctx := context.Background()
	print(testQueries)

	store := NewStore(dbpool)
	n := 4
	amount := int64(10)
	errs := make(chan error)
	results := make(chan TransferTxResult)

	accountOne, err := createRandomAccount(t)
	if err != nil {
		t.Error(err)
	}

	accountTwo, err := createRandomAccount(t)
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: accountOne.ID,
				ToAccountID:   accountTwo.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}

	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, accountOne.ID, transfer.FromAccountID)
		require.Equal(t, accountTwo.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(ctx, transfer.ID)
		require.NoError(t, err)

		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, accountOne.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(ctx, fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, accountTwo.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(ctx, toEntry.ID)
		require.NoError(t, err)

		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, accountOne.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, accountTwo.ID, toAccount.ID)

		diffOne := accountOne.Balance - fromAccount.Balance
		diffTwo := toAccount.Balance - accountTwo.Balance
		require.Equal(t, diffOne, diffTwo)
		require.True(t, diffOne > 0)
		require.True(t, diffOne%amount == 0)

		k := int(diffOne / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	updateAccountOne, err := testQueries.GetAccount(ctx, accountOne.ID)
	require.NoError(t, err)

	updateAccountTwo, err := testQueries.GetAccount(ctx, accountTwo.ID)
	require.NoError(t, err)

	require.Equal(t, accountOne.Balance-int64(n)*amount, updateAccountOne.Balance)
	require.Equal(t, accountTwo.Balance+int64(n)*amount, updateAccountTwo.Balance)
}
