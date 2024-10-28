package shoppingcart_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/ghouscht/shopping-cart/shoppingcart"
	"github.com/ghouscht/shopping-cart/shoppingcart/repo/mock"
)

func TestGetItems(t *testing.T) {
	tests := []struct {
		name            string
		setExpectations func(*mock.MockRepository)
		wantCode        int
		wantBody        string
	}{
		{
			name: "no error",
			setExpectations: func(mr *mock.MockRepository) {
				reservationID := 42
				mr.EXPECT().GetItems(gomock.Any(), 42).Return([]shoppingcart.Item{
					{
						Name:          "item 1",
						Quantity:      3,
						ReservationID: &reservationID,
					},
				}, nil)
			},
			wantCode: http.StatusOK,
			wantBody: `{"items":[{"name":"item 1", "quantity": 3, "reservationID": 42}]}`,
		},
		{
			name: "repo error",
			setExpectations: func(mr *mock.MockRepository) {
				mr.EXPECT().GetItems(gomock.Any(), 42).Return(nil, fmt.Errorf("boom"))
			},
			wantCode: http.StatusInternalServerError,
			wantBody: ``,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repo := mock.NewMockRepository(ctrl)

			mux := http.NewServeMux()
			shoppingcart.Register(mux, repo)

			tc.setExpectations(repo)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, "/items", nil)

			mux.ServeHTTP(recorder, request)
			require.Equal(t, tc.wantCode, recorder.Code)

			if tc.wantBody != "" {
				require.JSONEq(t, tc.wantBody, recorder.Body.String())
			}
		})
	}
}

func TestAddItem(t *testing.T) {
	tests := []struct {
		name            string
		setExpectations func(*mock.MockRepository)
		request         *http.Request
		wantCode        int
	}{
		{
			name: "ok",
			setExpectations: func(mr *mock.MockRepository) {
				mr.EXPECT().AddItem(gomock.Any(), 42, "item 1", 3).Return(nil)
			},
			request:  httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(`{"name": "item 1", "quantity": 3}`)),
			wantCode: http.StatusOK,
		},
		{
			name:            "empty name",
			setExpectations: func(mr *mock.MockRepository) {},
			request:         httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(`{"name": "", "quantity": 3}`)),
			wantCode:        http.StatusBadRequest,
		},
		{
			name:            "quantity <= 0",
			setExpectations: func(mr *mock.MockRepository) {},
			request:         httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(`{"name": "item 1", "quantity": -1}`)),
			wantCode:        http.StatusBadRequest,
		},
		{
			name:            "invalid json",
			setExpectations: func(mr *mock.MockRepository) {},
			request:         httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(`invalid`)),
			wantCode:        http.StatusBadRequest,
		},
		{
			name: "repo error",
			setExpectations: func(mr *mock.MockRepository) {
				mr.EXPECT().AddItem(gomock.Any(), 42, "item 1", 3).Return(fmt.Errorf("boom"))
			},
			request:  httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(`{"name": "item 1", "quantity": 3}`)),
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repo := mock.NewMockRepository(ctrl)

			mux := http.NewServeMux()
			shoppingcart.Register(mux, repo)

			tc.setExpectations(repo)

			recorder := httptest.NewRecorder()

			mux.ServeHTTP(recorder, tc.request)
			require.Equal(t, tc.wantCode, recorder.Code)
		})
	}
}
