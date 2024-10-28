package shoppingcart

import (
	"context"
	"log/slog"
	"time"
)

// NewReservationProcessor reads unreserved items from the database, then tries to reserve them, if this is successful
// the item will be marked as reserved in the database. This function blocks until the given context is cancelled. Each
// reservation is handled sequentially, this can be a problem as the external service is slow and we might not be able
// to cope with the amount of reservations we should do.
// TODO: Implement a worker pool to process the reservations in a more parallel manner.
func NewReservationProcessor(ctx context.Context, interval *time.Ticker, repo Repository, reservation Reservation) {
	slog.Info("starting reservation processor")

	for {
		select {
		case <-ctx.Done():
			slog.Info("stopping reservation processor")
			return
		case <-interval.C:
			slog.Debug("processing pending reservations")
		}

		items, err := repo.GetUnreserved(ctx)
		if err != nil {
			slog.Error("reading unreserved items", slog.Any("err", err))
			continue
		}

		for _, item := range items {
			reserveCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
			reservationID, err := reservation.ReserveItem(reserveCtx, item.Name, item.Quantity)
			if err != nil {
				cancel()
				slog.Error("reserving item", slog.Any("err", err), slog.String("item", item.Name))
				continue
			}

			if err := repo.MarkReserved(reserveCtx, item.UserID, item.Name, reservationID); err != nil {
				cancel()
				// TODO: We reserved the item but failed to persist. Ideally we should delete the reservation now.
				slog.Error("mark item as reserved", slog.Any("err", err), slog.String("item", item.Name))
				continue
			}
			cancel()
		}
	}
}
