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
