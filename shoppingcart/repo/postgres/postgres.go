package postgres

import (
	"context"
	"database/sql"
	_ "embed"
)

type ShoppingCartRepository struct {
	db *sql.DB
}

//go:embed schema.sql
var schema string

func NewShoppingCartRepository(db *sql.DB) ShoppingCartRepository {
	// Create the database schema. Usually this should be done by some sort of database migration tooling. To keep things
	// a bit simple for the sake of this exercise I simply do this here when we create the ShoppingCartRepository struct.
	_, err := db.ExecContext(context.Background(), schema)
	if err != nil {
		panic(err)
	}

	return ShoppingCartRepository{db: db}
}
