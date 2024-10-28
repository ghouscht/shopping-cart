package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/ghouscht/shopping-cart/shoppingcart"
	"github.com/ghouscht/shopping-cart/shoppingcart/repo/postgres"
	"github.com/ghouscht/shopping-cart/shoppingcart/reservation/dummy"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	db, err := sql.Open("pgx", "postgres://shoppingcart:V3ry$3cr3t@localhost:5432/shoppingcart?sslmode=disable")
	if err != nil {
		slog.Error("open db", slog.Any("err", err))
		os.Exit(1)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		slog.Error("ping database", slog.Any("err", err))
		os.Exit(1)
	}

	shoppingcartRepo := postgres.NewShoppingCartRepository(db)
	reservationService := dummy.ReservationService{}

	if err := run(ctx, shoppingcartRepo, reservationService); err != nil {
		slog.Error("run shopping cart", slog.Any("err", err))
		os.Exit(1) // TODO: skips defer
	}
}

func run(ctx context.Context, repo shoppingcart.Repository, reservation shoppingcart.Reservation) error {
	var wg sync.WaitGroup

	const serverAddr = ":8080"

	mux := http.NewServeMux()
	shoppingcart.Register(mux, repo)

	server := http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}

	shutdownErr := make(chan error)

	wg.Add(1)
	go func() {
		defer wg.Done()
		shoppingcart.NewReservationProcessor(ctx, time.NewTicker(5*time.Second), repo, reservation)
	}()

	go func() {
		<-ctx.Done() // block until the incoming context is cancelled
		slog.Info("stopping http server")

		// Create a context to timeout the shutdown operation without the using the incoming context. When this code is
		// reached the incoming context is already cancelled.
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			shutdownErr <- fmt.Errorf("http shutdown: %w", err)
		}

		close(shutdownErr)
	}()

	slog.Info("starting http server", slog.String("addr", serverAddr))

	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("http listen: %w", err)
	}

	wg.Wait()

	return <-shutdownErr
}
