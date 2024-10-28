package shoppingcart

import "context"

//go:generate go run go.uber.org/mock/mockgen -typed -source ports.go -destination repo/mock/mocks_gen.go -package mock
type Repository interface {
	GetItems(ctx context.Context, userID int) ([]Item, error)
	AddItem(ctx context.Context, userID int, name string, quantity int) error
	MarkReserved(ctx context.Context, userID int, name string, reservationID int) error
	GetUnreserved(ctx context.Context) ([]Item, error)
}

//go:generate go run go.uber.org/mock/mockgen -typed -source ports.go -destination reservation/mock/mocks_gen.go -package mock
type Reservation interface {
	ReserveItem(ctx context.Context, item string, quantity int) (int, error)
}

// Item represents a single item in the shopping cart. For the sake of simplicity the Repository uses the same data
// type to avoid the need for a conversion. Usually the repository should use a distinct type.
type Item struct {
	Name          string `json:"name"`
	UserID        int    `json:"userID,omitempty"`
	Quantity      int    `json:"quantity"`
	ReservationID *int   `json:"reservationID"`
}
