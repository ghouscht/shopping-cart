// Code generated by MockGen. DO NOT EDIT.
// Source: ports.go
//
// Generated by this command:
//
//	mockgen -typed -source ports.go -destination repo/mock/mocks_gen.go -package mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	shoppingcart "github.com/ghouscht/shopping-cart/shoppingcart"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
	isgomock struct{}
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddItem mocks base method.
func (m *MockRepository) AddItem(ctx context.Context, userID int, name string, quantity int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddItem", ctx, userID, name, quantity)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddItem indicates an expected call of AddItem.
func (mr *MockRepositoryMockRecorder) AddItem(ctx, userID, name, quantity any) *MockRepositoryAddItemCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddItem", reflect.TypeOf((*MockRepository)(nil).AddItem), ctx, userID, name, quantity)
	return &MockRepositoryAddItemCall{Call: call}
}

// MockRepositoryAddItemCall wrap *gomock.Call
type MockRepositoryAddItemCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryAddItemCall) Return(arg0 error) *MockRepositoryAddItemCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryAddItemCall) Do(f func(context.Context, int, string, int) error) *MockRepositoryAddItemCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryAddItemCall) DoAndReturn(f func(context.Context, int, string, int) error) *MockRepositoryAddItemCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetItems mocks base method.
func (m *MockRepository) GetItems(ctx context.Context, userID int) ([]shoppingcart.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItems", ctx, userID)
	ret0, _ := ret[0].([]shoppingcart.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItems indicates an expected call of GetItems.
func (mr *MockRepositoryMockRecorder) GetItems(ctx, userID any) *MockRepositoryGetItemsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockRepository)(nil).GetItems), ctx, userID)
	return &MockRepositoryGetItemsCall{Call: call}
}

// MockRepositoryGetItemsCall wrap *gomock.Call
type MockRepositoryGetItemsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryGetItemsCall) Return(arg0 []shoppingcart.Item, arg1 error) *MockRepositoryGetItemsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryGetItemsCall) Do(f func(context.Context, int) ([]shoppingcart.Item, error)) *MockRepositoryGetItemsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryGetItemsCall) DoAndReturn(f func(context.Context, int) ([]shoppingcart.Item, error)) *MockRepositoryGetItemsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetUnreserved mocks base method.
func (m *MockRepository) GetUnreserved(ctx context.Context) ([]shoppingcart.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnreserved", ctx)
	ret0, _ := ret[0].([]shoppingcart.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnreserved indicates an expected call of GetUnreserved.
func (mr *MockRepositoryMockRecorder) GetUnreserved(ctx any) *MockRepositoryGetUnreservedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnreserved", reflect.TypeOf((*MockRepository)(nil).GetUnreserved), ctx)
	return &MockRepositoryGetUnreservedCall{Call: call}
}

// MockRepositoryGetUnreservedCall wrap *gomock.Call
type MockRepositoryGetUnreservedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryGetUnreservedCall) Return(arg0 []shoppingcart.Item, arg1 error) *MockRepositoryGetUnreservedCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryGetUnreservedCall) Do(f func(context.Context) ([]shoppingcart.Item, error)) *MockRepositoryGetUnreservedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryGetUnreservedCall) DoAndReturn(f func(context.Context) ([]shoppingcart.Item, error)) *MockRepositoryGetUnreservedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MarkReserved mocks base method.
func (m *MockRepository) MarkReserved(ctx context.Context, userID int, name string, reservationID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkReserved", ctx, userID, name, reservationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkReserved indicates an expected call of MarkReserved.
func (mr *MockRepositoryMockRecorder) MarkReserved(ctx, userID, name, reservationID any) *MockRepositoryMarkReservedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkReserved", reflect.TypeOf((*MockRepository)(nil).MarkReserved), ctx, userID, name, reservationID)
	return &MockRepositoryMarkReservedCall{Call: call}
}

// MockRepositoryMarkReservedCall wrap *gomock.Call
type MockRepositoryMarkReservedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryMarkReservedCall) Return(arg0 error) *MockRepositoryMarkReservedCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryMarkReservedCall) Do(f func(context.Context, int, string, int) error) *MockRepositoryMarkReservedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryMarkReservedCall) DoAndReturn(f func(context.Context, int, string, int) error) *MockRepositoryMarkReservedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockReservation is a mock of Reservation interface.
type MockReservation struct {
	ctrl     *gomock.Controller
	recorder *MockReservationMockRecorder
	isgomock struct{}
}

// MockReservationMockRecorder is the mock recorder for MockReservation.
type MockReservationMockRecorder struct {
	mock *MockReservation
}

// NewMockReservation creates a new mock instance.
func NewMockReservation(ctrl *gomock.Controller) *MockReservation {
	mock := &MockReservation{ctrl: ctrl}
	mock.recorder = &MockReservationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReservation) EXPECT() *MockReservationMockRecorder {
	return m.recorder
}

// ReserveItem mocks base method.
func (m *MockReservation) ReserveItem(ctx context.Context, item string, quantity int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReserveItem", ctx, item, quantity)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReserveItem indicates an expected call of ReserveItem.
func (mr *MockReservationMockRecorder) ReserveItem(ctx, item, quantity any) *MockReservationReserveItemCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReserveItem", reflect.TypeOf((*MockReservation)(nil).ReserveItem), ctx, item, quantity)
	return &MockReservationReserveItemCall{Call: call}
}

// MockReservationReserveItemCall wrap *gomock.Call
type MockReservationReserveItemCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockReservationReserveItemCall) Return(arg0 int, arg1 error) *MockReservationReserveItemCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockReservationReserveItemCall) Do(f func(context.Context, string, int) (int, error)) *MockReservationReserveItemCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockReservationReserveItemCall) DoAndReturn(f func(context.Context, string, int) (int, error)) *MockReservationReserveItemCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
