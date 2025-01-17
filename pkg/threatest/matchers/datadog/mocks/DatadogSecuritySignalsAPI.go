// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	v2datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	mock "github.com/stretchr/testify/mock"
)

// DatadogSecuritySignalsAPI is an autogenerated mock type for the DatadogSecuritySignalsAPI type
type DatadogSecuritySignalsAPI struct {
	mock.Mock
}

// CloseSignal provides a mock function with given fields: id
func (_m *DatadogSecuritySignalsAPI) CloseSignal(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchSignals provides a mock function with given fields: query
func (_m *DatadogSecuritySignalsAPI) SearchSignals(query string) ([]v2datadog.SecurityMonitoringSignal, error) {
	ret := _m.Called(query)

	var r0 []v2datadog.SecurityMonitoringSignal
	if rf, ok := ret.Get(0).(func(string) []v2datadog.SecurityMonitoringSignal); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v2datadog.SecurityMonitoringSignal)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDatadogSecuritySignalsAPI interface {
	mock.TestingT
	Cleanup(func())
}

// NewDatadogSecuritySignalsAPI creates a new instance of DatadogSecuritySignalsAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDatadogSecuritySignalsAPI(t mockConstructorTestingTNewDatadogSecuritySignalsAPI) *DatadogSecuritySignalsAPI {
	mock := &DatadogSecuritySignalsAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
