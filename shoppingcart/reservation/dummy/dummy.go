// Package dummy implements a dummy ReservationService that simply returns a random number as a reservation ID.
package dummy

import (
	"context"
	"math/rand/v2"
)

type ReservationService struct{}

func (ReservationService) ReserveItem(_ context.Context, _ string, _ int) (int, error) {
	return rand.IntN(1000), nil
}
