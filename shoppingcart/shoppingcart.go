package shoppingcart

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// Item represents a single item in the shopping cart. For the sake of simplicity the Repository uses the same data
// type to avoid the need for a conversion. Usually the repository should use a distinct type.
type Item struct {
	Name          string `json:"name"`
	UserID        int    `json:"userID,omitempty"`
	Quantity      int    `json:"quantity"`
	ReservationID *int   `json:"reservationID"`
}

type GetItemsResponse struct {
	Items []Item `json:"items"`
}

type AddItemRequest struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

// Register registers the shoppingcart routes to the given ServeMux.
func Register(mux *http.ServeMux, repo Repository) {
	mux.HandleFunc("GET /items", getItems(repo))
	mux.HandleFunc("POST /items", addItem(repo))
}

// The application currently only handles a single users shopping cart. In a real application there would probably be
// shopping carts for many users. In such a case I would expect that there is some kind of authentication middleware
// that e.g. places a userID in a `context.Context` object which can be retrieved from the handler functions and then
// used to store the shopping cart. For the sake of simplicity of this task the userID will always be 42.
const userID = 42

func getItems(repo Repository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := repo.GetItems(r.Context(), userID)
		if err != nil {
			slog.Error("read items from repository", slog.Any("err", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(GetItemsResponse{Items: items}); err != nil {
			slog.Error("encode json", slog.Any("err", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func addItem(repo Repository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestData AddItemRequest

		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			slog.Error("decode request", slog.Any("err", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if requestData.Name == "" || requestData.Quantity <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := repo.AddItem(r.Context(), userID, requestData.Name, requestData.Quantity); err != nil {
			slog.Error("add item to repository", slog.Any("err", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
