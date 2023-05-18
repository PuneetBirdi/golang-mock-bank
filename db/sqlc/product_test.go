package db

import (
	"context"
	"testing"

	"github.com/PuneetBirdi/golang-bank/util"
	"github.com/stretchr/testify/require"
)

func getRandomProduct(t *testing.T) Product {
	arg := util.RandomInt(1, 4)

	product, err := testQueries.GetProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	return product
}

