package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddGetItem(t *testing.T) {
	err := repo.AddItem(context.Background(), 42, "item 1", 3)
	require.NoError(t, err)

	// second add should update the quantity
	err = repo.AddItem(context.Background(), 42, "item 1", 1)
	require.NoError(t, err)

	// read back and compare
	items, err := repo.GetItems(context.Background(), 42)
	require.NoError(t, err)
	require.Len(t, items, 1)
	require.Equal(t, "item 1", items[0].Name)
	require.Equal(t, 4, items[0].Quantity)
	require.Empty(t, items[0].ReservationID)
}
