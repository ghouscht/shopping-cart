package postgres_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/ory/dockertest/v3"

	"github.com/ghouscht/shopping-cart/shoppingcart/repo/postgres"
)

var repo postgres.ShoppingCartRepository

// Start a docker container running postgres to execute the tests against a real database.
// TODO: The repo and thus the state is shared between all tests which is an anti pattern and could easily cause trouble.
// For the sake of simplicity I intentionally didn't spend more time on this. I think you get the point about how I would
// test a package which interacts with a database.
func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("postgres", "16.4-alpine", []string{
		"POSTGRES_PASSWORD=password",
		"POSTGRES_USER=shoppingcart",
		"POSTGRES_DB=shoppingcart",
		"listen_addresses = '*'",
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	var db *sql.DB
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("pgx", fmt.Sprintf("postgres://shoppingcart:password@localhost:%s/shoppingcart?sslmode=disable", resource.GetPort("5432/tcp")))
		if err != nil {
			return err
		}

		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	repo = postgres.NewShoppingCartRepository(db)

	// run the tests
	result := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(result)
}
