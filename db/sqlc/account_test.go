package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/louischering/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) (Account, error) {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account, err
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	accountOne, _ := createRandomAccount(t)
	accountTwo, err := testQueries.GetAccount(context.Background(), accountOne.ID)

	require.NoError(t, err)
	require.NotEmpty(t, accountTwo)
	require.Equal(t, accountOne.ID, accountTwo.ID)
	require.Equal(t, accountOne.CreatedAt, accountTwo.CreatedAt)
	require.Equal(t, accountOne.Balance, accountTwo.Balance)
	require.Equal(t, accountOne.Owner, accountTwo.Owner)
}

func TestDeleteAccount(t *testing.T) {
	accountOne, _ := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), accountOne.ID)

	require.NoError(t, err)
	_, err = testQueries.GetAccount(context.Background(), accountOne.ID)
	require.Error(t, sql.ErrNoRows, err)
}

func TestListAccounts(t *testing.T) {
	accountSeedAmount := 5
	for i := 0; i <= accountSeedAmount; i++ {
		createRandomAccount(t)
	}
	params := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	require.Equal(t, len(accounts), accountSeedAmount)
}
