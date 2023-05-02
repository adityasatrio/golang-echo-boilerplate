// Code generated by mockery v2.20.0. DO NOT EDIT.

package mock_transaction

import (
	context "context"
	ent "myapp/ent"

	mock "github.com/stretchr/testify/mock"
)

// TrxService is an autogenerated mock type for the TrxService type
type TrxService struct {
	mock.Mock
}

// WithTx provides a mock function with given fields: ctx, fn
func (_m *TrxService) WithTx(ctx context.Context, fn func(*ent.Tx) error) error {
	ret := _m.Called(ctx, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(*ent.Tx) error) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTrxService interface {
	mock.TestingT
	Cleanup(func())
}

// NewTrxService creates a new instance of TrxService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTrxService(t mockConstructorTestingTNewTrxService) *TrxService {
	mock := &TrxService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}