package db

import (
	"context"
	"testing"
	"github.com/PuneetBirdi/golang-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		FullName: (util.RandomString(5) + " " + util.RandomString(8)),
		HashedPassword: util.RandomString(16),
		Email: (util.RandomString(5) + "@gmail.com"),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.ID)

	return user
}
