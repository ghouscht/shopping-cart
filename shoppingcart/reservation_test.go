package shoppingcart_test

import (
	"context"
	"testing"
	"time"

	"go.uber.org/mock/gomock"

	"github.com/ghouscht/shopping-cart/shoppingcart"
	"github.com/ghouscht/shopping-cart/shoppingcart/repo/mock"
	mock_reservation "github.com/ghouscht/shopping-cart/shoppingcart/reservation/mock"
)

// TODO: There should be some more test cases for the error branches but I think this is enough so you get the point how
// I would test this part of the code.
func TestReservationProcessor(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock.NewMockRepository(ctrl)
	reservationService := mock_reservation.NewMockReservation(ctrl)

	ctx, cancel := context.WithCancel(context.Background())

	repo.EXPECT().GetUnreserved(gomock.Any()).DoAndReturn(func(ctx context.Context) ([]shoppingcart.Item, error) {
		// Cancel the context before returning the two unreserved items to make sure the ReservationProcessor exits
		// after processing the reservations.
		cancel()

		return []shoppingcart.Item{
			{
				Name:     "item 1",
				UserID:   1,
				Quantity: 1,
			},
			{
				Name:     "item 2",
				UserID:   2,
				Quantity: 2,
			},
		}, nil
	})

	// item 1
	reservationService.EXPECT().ReserveItem(gomock.Any(), "item 1", 1).Return(1111, nil)
	repo.EXPECT().MarkReserved(gomock.Any(), 1, "item 1", 1111).Return(nil)

	// item 2
	reservationService.EXPECT().ReserveItem(gomock.Any(), "item 2", 2).Return(2222, nil)
	repo.EXPECT().MarkReserved(gomock.Any(), 2, "item 2", 2222).Return(nil)

	shoppingcart.NewReservationProcessor(ctx, time.NewTicker(1*time.Second), repo, reservationService)
}
