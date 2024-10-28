package postgres

import (
	"context"
	"fmt"

	"github.com/ghouscht/shopping-cart/shoppingcart"
)

const markReservedSQL = "UPDATE items SET reservation_id = $1 WHERE user_id = $2 AND name = $3"

func (r ShoppingCartRepository) MarkReserved(ctx context.Context, userID int, name string, reservationID int) error {
	_, err := r.db.ExecContext(ctx, markReservedSQL, reservationID, userID, name)
	if err != nil {
		return fmt.Errorf("mark item as reserved: %w", err)
	}

	return nil
}

const getUnreservedSQL = "SELECT name, user_id, quantity FROM items WHERE reservation_id IS NULL"

func (r ShoppingCartRepository) GetUnreserved(ctx context.Context) ([]shoppingcart.Item, error) {
	rows, err := r.db.QueryContext(ctx, getUnreservedSQL)
	if err != nil {
		return nil, fmt.Errorf("query unreserved items: %w", err)
	}

	var items []shoppingcart.Item

	for rows.Next() {
		var item shoppingcart.Item

		err := rows.Scan(&item.Name, &item.UserID, &item.Quantity)
		if err != nil {
			return nil, fmt.Errorf("scan item row: %w", err)
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating unreserved item rows: %w", err)
	}

	return items, nil
}
