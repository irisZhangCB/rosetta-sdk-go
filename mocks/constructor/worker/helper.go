// Code generated by mockery v1.0.0. DO NOT EDIT.

package worker

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	keys "github.com/irisZhangCB/rosetta-sdk-go/keys"
	database "github.com/irisZhangCB/rosetta-sdk-go/storage/database"
	types "github.com/irisZhangCB/rosetta-sdk-go/types"
)

// Helper is an autogenerated mock type for the Helper type
type Helper struct {
	mock.Mock
}

// AllAccounts provides a mock function with given fields: _a0, _a1
func (_m *Helper) AllAccounts(_a0 context.Context, _a1 database.Transaction) ([]*types.AccountIdentifier, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*types.AccountIdentifier
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction) []*types.AccountIdentifier); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.AccountIdentifier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Transaction) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Balance provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Helper) Balance(_a0 context.Context, _a1 database.Transaction, _a2 *types.AccountIdentifier, _a3 *types.Currency) (*types.Amount, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *types.Amount
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, *types.AccountIdentifier, *types.Currency) *types.Amount); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Amount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Transaction, *types.AccountIdentifier, *types.Currency) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Coins provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Helper) Coins(_a0 context.Context, _a1 database.Transaction, _a2 *types.AccountIdentifier, _a3 *types.Currency) ([]*types.Coin, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 []*types.Coin
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, *types.AccountIdentifier, *types.Currency) []*types.Coin); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Coin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Transaction, *types.AccountIdentifier, *types.Currency) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Derive provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Helper) Derive(_a0 context.Context, _a1 *types.NetworkIdentifier, _a2 *types.PublicKey, _a3 map[string]interface{}) (*types.AccountIdentifier, map[string]interface{}, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *types.AccountIdentifier
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, *types.PublicKey, map[string]interface{}) *types.AccountIdentifier); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.AccountIdentifier)
		}
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, *types.PublicKey, map[string]interface{}) map[string]interface{}); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *types.NetworkIdentifier, *types.PublicKey, map[string]interface{}) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetBlob provides a mock function with given fields: ctx, dbTx, key
func (_m *Helper) GetBlob(ctx context.Context, dbTx database.Transaction, key string) (bool, []byte, error) {
	ret := _m.Called(ctx, dbTx, key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, string) bool); ok {
		r0 = rf(ctx, dbTx, key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(context.Context, database.Transaction, string) []byte); ok {
		r1 = rf(ctx, dbTx, key)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, database.Transaction, string) error); ok {
		r2 = rf(ctx, dbTx, key)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// LockedAccounts provides a mock function with given fields: _a0, _a1
func (_m *Helper) LockedAccounts(_a0 context.Context, _a1 database.Transaction) ([]*types.AccountIdentifier, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*types.AccountIdentifier
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction) []*types.AccountIdentifier); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.AccountIdentifier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Transaction) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetBlob provides a mock function with given fields: ctx, dbTx, key, value
func (_m *Helper) SetBlob(ctx context.Context, dbTx database.Transaction, key string, value []byte) error {
	ret := _m.Called(ctx, dbTx, key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, string, []byte) error); ok {
		r0 = rf(ctx, dbTx, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreKey provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Helper) StoreKey(_a0 context.Context, _a1 database.Transaction, _a2 *types.AccountIdentifier, _a3 *keys.KeyPair) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, *types.AccountIdentifier, *keys.KeyPair) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
