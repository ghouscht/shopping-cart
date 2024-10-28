package http_test

import (
	"context"
	stdhttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ghouscht/shopping-cart/shoppingcart/reservation/http"
)

func TestReserveItem(t *testing.T) {
	testServer := httptest.NewServer(stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		_, _ = w.Write([]byte(`{"reservation_id": 1234}`))
	}))
	defer testServer.Close()

	svc := http.NewReservationService(testServer.URL)

	reservationID, err := svc.ReserveItem(context.Background(), "item 1", 3)
	require.NoError(t, err)
	require.Equal(t, 1234, reservationID)
}
