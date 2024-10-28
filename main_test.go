package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/ghouscht/shopping-cart/shoppingcart/repo/mock"
	mock_reservation "github.com/ghouscht/shopping-cart/shoppingcart/reservation/mock"
)

// TODO: This test doesn't really verify anything but serves as an example how I could test the logic from main. The
// main function is only used to perform setup (e.g. of database) and then control is handed over to `run` which does
// the actual magic.
func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock.NewMockRepository(ctrl)
	reservationService := mock_reservation.NewMockReservation(ctrl)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := run(ctx, repo, reservationService)
	require.NoError(t, err)
}
