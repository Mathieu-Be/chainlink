// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	eth "chainlink/core/services/eth"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"
)

// LogBroadcaster is an autogenerated mock type for the LogBroadcaster type
type LogBroadcaster struct {
	mock.Mock
}

// Register provides a mock function with given fields: address, listener
func (_m *LogBroadcaster) Register(address common.Address, listener eth.LogListener) {
	_m.Called(address, listener)
}

// Start provides a mock function with given fields:
func (_m *LogBroadcaster) Start() {
	_m.Called()
}

// Stop provides a mock function with given fields:
func (_m *LogBroadcaster) Stop() {
	_m.Called()
}

// Unregister provides a mock function with given fields: address, listener
func (_m *LogBroadcaster) Unregister(address common.Address, listener eth.LogListener) {
	_m.Called(address, listener)
}
