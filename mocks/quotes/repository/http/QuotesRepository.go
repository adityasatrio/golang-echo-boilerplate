// Code generated by mockery v2.30.1. DO NOT EDIT.

package mock_http

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// QuotesRepository is an autogenerated mock type for the QuotesRepository type
type QuotesRepository struct {
	mock.Mock
}

// GetQuotes provides a mock function with given fields: ctx
func (_m *QuotesRepository) GetQuotes(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewQuotesRepository creates a new instance of QuotesRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuotesRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *QuotesRepository {
	mock := &QuotesRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}