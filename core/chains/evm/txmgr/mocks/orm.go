// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	mock "github.com/stretchr/testify/mock"

	txmgr "github.com/smartcontractkit/chainlink/core/chains/evm/txmgr"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// EthTransactions provides a mock function with given fields: offset, limit
func (_m *ORM) EthTransactions(offset int, limit int) ([]txmgr.EthTx, int, error) {
	ret := _m.Called(offset, limit)

	var r0 []txmgr.EthTx
	if rf, ok := ret.Get(0).(func(int, int) []txmgr.EthTx); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgr.EthTx)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EthTransactionsWithAttempts provides a mock function with given fields: offset, limit
func (_m *ORM) EthTransactionsWithAttempts(offset int, limit int) ([]txmgr.EthTx, int, error) {
	ret := _m.Called(offset, limit)

	var r0 []txmgr.EthTx
	if rf, ok := ret.Get(0).(func(int, int) []txmgr.EthTx); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgr.EthTx)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EthTxAttempts provides a mock function with given fields: offset, limit
func (_m *ORM) EthTxAttempts(offset int, limit int) ([]txmgr.EthTxAttempt, int, error) {
	ret := _m.Called(offset, limit)

	var r0 []txmgr.EthTxAttempt
	if rf, ok := ret.Get(0).(func(int, int) []txmgr.EthTxAttempt); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgr.EthTxAttempt)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindEthTxAttempt provides a mock function with given fields: hash
func (_m *ORM) FindEthTxAttempt(hash common.Hash) (*txmgr.EthTxAttempt, error) {
	ret := _m.Called(hash)

	var r0 *txmgr.EthTxAttempt
	if rf, ok := ret.Get(0).(func(common.Hash) *txmgr.EthTxAttempt); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*txmgr.EthTxAttempt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxAttemptsByEthTxIDs provides a mock function with given fields: ids
func (_m *ORM) FindEthTxAttemptsByEthTxIDs(ids []int64) ([]txmgr.EthTxAttempt, error) {
	ret := _m.Called(ids)

	var r0 []txmgr.EthTxAttempt
	if rf, ok := ret.Get(0).(func([]int64) []txmgr.EthTxAttempt); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgr.EthTxAttempt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]int64) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxByHash provides a mock function with given fields: hash
func (_m *ORM) FindEthTxByHash(hash common.Hash) (*txmgr.EthTx, error) {
	ret := _m.Called(hash)

	var r0 *txmgr.EthTx
	if rf, ok := ret.Get(0).(func(common.Hash) *txmgr.EthTx); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*txmgr.EthTx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxWithAttempts provides a mock function with given fields: etxID
func (_m *ORM) FindEthTxWithAttempts(etxID int64) (txmgr.EthTx, error) {
	ret := _m.Called(etxID)

	var r0 txmgr.EthTx
	if rf, ok := ret.Get(0).(func(int64) txmgr.EthTx); ok {
		r0 = rf(etxID)
	} else {
		r0 = ret.Get(0).(txmgr.EthTx)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(etxID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertEthReceipt provides a mock function with given fields: receipt
func (_m *ORM) InsertEthReceipt(receipt *txmgr.EthReceipt) error {
	ret := _m.Called(receipt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*txmgr.EthReceipt) error); ok {
		r0 = rf(receipt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertEthTx provides a mock function with given fields: etx
func (_m *ORM) InsertEthTx(etx *txmgr.EthTx) error {
	ret := _m.Called(etx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*txmgr.EthTx) error); ok {
		r0 = rf(etx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertEthTxAttempt provides a mock function with given fields: attempt
func (_m *ORM) InsertEthTxAttempt(attempt *txmgr.EthTxAttempt) error {
	ret := _m.Called(attempt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*txmgr.EthTxAttempt) error); ok {
		r0 = rf(attempt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewORM interface {
	mock.TestingT
	Cleanup(func())
}

// NewORM creates a new instance of ORM. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewORM(t mockConstructorTestingTNewORM) *ORM {
	mock := &ORM{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
