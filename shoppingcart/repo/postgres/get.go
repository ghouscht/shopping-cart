package postgres

import (
	"context"
	"fmt"

	"github.com/ghouscht/shopping-cart/shoppingcart"
)

const getSQL = "SELECT name, quantity, reservation_id FROM items WHERE user_id = $1"

func (r ShoppingCartRepository) GetItems(ctx context.Context, userID int) ([]shoppingcart.Item, error) {
	rows, err := r.db.QueryContext(ctx, getSQL, userID)
	if err != nil {
		return nil, fmt.Errorf("query items: %w", err)
	}

	var items []shoppingcart.Item

	for rows.Next() {
		var item shoppingcart.Item

		err := rows.Scan(&item.Name, &item.Quantity, &item.ReservationID)
		if err != nil {
			return nil, fmt.Errorf("scan item row: %w", err)
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating rows: %w", err)
	}

	return items, nil
}
