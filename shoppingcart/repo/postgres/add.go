package postgres

import (
	"context"
	"fmt"
)

const addSQL = `
INSERT INTO items (user_id, name, quantity)
VALUES ($1, $2, $3)
ON CONFLICT ON CONSTRAINT items_user_id_name_key DO UPDATE
SET quantity = items.quantity + $3,
reservation_id = NULL
`

// AddItem adds and item to the shopping cart. If the item is already in the shopping cart the quantity is increased by
// the requested amount and if a reservation was made it is deleted.
// TODO: Deleting the reservation in the database is not ideal as we won't tell the reservation system that this reservation
// is no longer valid. However for the sake of simplicity I think this should be fine.
func (r ShoppingCartRepository) AddItem(ctx context.Context, userID int, name string, quantity int) error {
	_, err := r.db.ExecContext(ctx, addSQL, userID, name, quantity)
	if err != nil {
		return fmt.Errorf("insert item: %w", err)
	}

	return nil
}
