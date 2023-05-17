package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/PuneetBirdi/golang-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.ID,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
		ProductType: ProductsEnum(util.RandomProductType()),
		AccountType: AccountsEnum(util.RandomAccountType()),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	args := UpdateAccountParams {
		ID: account.ID,
		Balance: util.RandomMoney(),
	}
	
	updatedAccount, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)

	require.Equal(t, updatedAccount.ID, args.ID)
	require.Equal(t, updatedAccount.Balance, args.Balance)
	require.Equal(t, updatedAccount.Owner, account.Owner)
	require.Equal(t, updatedAccount.Currency, account.Currency)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	args := account.ID

	err := testQueries.DeleteAccount(context.Background(), args)
	require.NoError(t, err)

	emptyResult, err := testQueries.GetAccount(context.Background(), args)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, emptyResult)
}

func TestListAccount(t *testing.T){
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams {
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
