package dummy_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ghouscht/shopping-cart/shoppingcart/reservation/dummy"
)

func TestReserveItem(t *testing.T) {
	svc := dummy.ReservationService{}

	reservationID, err := svc.ReserveItem(context.Background(), "", 0)
	require.NoError(t, err)
	require.GreaterOrEqual(t, reservationID, 0)
}
