// Code generated by mockery v1.0.0. DO NOT EDIT.

package coordinator

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/irisZhangCB/rosetta-sdk-go/types"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// TransactionCreated provides a mock function with given fields: _a0, _a1, _a2
func (_m *Handler) TransactionCreated(_a0 context.Context, _a1 string, _a2 *types.TransactionIdentifier) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *types.TransactionIdentifier) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
