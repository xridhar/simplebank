package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xridhar/simplebank/util"
)


func createRandomEntry(t *testing.T) Entry {
	arg := CreateEntryParams{
		AccountID: util.RandomInt(1, 8),
		Amount: util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreateAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}