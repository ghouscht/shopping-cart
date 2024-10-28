package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReservations(t *testing.T) {
	err := repo.AddItem(context.Background(), 1, "item 1", 3)
	require.NoError(t, err)

	err = repo.AddItem(context.Background(), 2, "item 2", 1)
	require.NoError(t, err)

	items, err := repo.GetUnreserved(context.Background())
	require.NoError(t, err)

	for _, item := range items {
		require.Nil(t, item.ReservationID)
	}

	err = repo.MarkReserved(context.Background(), 1, "item 1", 123)
	require.NoError(t, err)

	items, err = repo.GetItems(context.Background(), 1)
	require.NoError(t, err)
	require.Len(t, items, 1)
	require.NotNil(t, items[0].ReservationID)
	require.Equal(t, 123, *items[0].ReservationID)
}
