// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import internal "github.com/kyma-project/control-plane/components/kyma-environment-broker/internal"
import mock "github.com/stretchr/testify/mock"

import time "time"

// Strategy is an autogenerated mock type for the Strategy type
type Strategy struct {
	mock.Mock
}

// Execute provides a mock function with given fields: operations
func (_m *Strategy) Execute(operations []internal.RuntimeOperation) (time.Duration, error) {
	ret := _m.Called(operations)

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func([]internal.RuntimeOperation) time.Duration); ok {
		r0 = rf(operations)
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]internal.RuntimeOperation) error); ok {
		r1 = rf(operations)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}