// Package http implements the ReservationService through an http endpoint. This only serves as an example because this
// endpoint does not exist.
package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ReservationService implements the Service interface and would call an external (http) service to reserve an item.
type ReservationService struct {
	client *http.Client
	url    string
}

func NewReservationService(url string) ReservationService {
	return ReservationService{client: &http.Client{}, url: url}
}

type ReserveReq struct {
	Item     string `json:"item"`
	Quantity int    `json:"quantity"`
}

type ReserveResp struct {
	ReservationID int `json:"reservation_id"`
}

func (s *ReservationService) ReserveItem(ctx context.Context, item string, quantity int) (int, error) {
	data, err := json.Marshal(ReserveReq{Item: item, Quantity: quantity})
	if err != nil {
		return 0, fmt.Errorf("marshal reserve request: %w", err)
	}

	resp, err := s.client.Post(s.url+"/reserve", "application/json", bytes.NewReader(data))
	if err != nil {
		return 0, fmt.Errorf("reserve request: %w", err)
	}
	defer resp.Body.Close()

	var reserveResp ReserveResp
	err = json.NewDecoder(resp.Body).Decode(&reserveResp)
	if err != nil {
		return 0, fmt.Errorf("decoding response: %w", err)
	}

	return reserveResp.ReservationID, nil
}
