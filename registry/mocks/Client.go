// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	types "github.com/denniskertis/go-mod-registry/v42/pkg/types"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// GetAllServiceEndpoints provides a mock function with given fields:
func (_m *Client) GetAllServiceEndpoints() ([]types.ServiceEndpoint, error) {
	ret := _m.Called()

	var r0 []types.ServiceEndpoint
	if rf, ok := ret.Get(0).(func() []types.ServiceEndpoint); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.ServiceEndpoint)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServiceEndpoint provides a mock function with given fields: serviceId
func (_m *Client) GetServiceEndpoint(serviceId string) (types.ServiceEndpoint, error) {
	ret := _m.Called(serviceId)

	var r0 types.ServiceEndpoint
	if rf, ok := ret.Get(0).(func(string) types.ServiceEndpoint); ok {
		r0 = rf(serviceId)
	} else {
		r0 = ret.Get(0).(types.ServiceEndpoint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(serviceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsAlive provides a mock function with given fields:
func (_m *Client) IsAlive() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsServiceAvailable provides a mock function with given fields: serviceId
func (_m *Client) IsServiceAvailable(serviceId string) (bool, error) {
	ret := _m.Called(serviceId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(serviceId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(serviceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields:
func (_m *Client) Register() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterCheck provides a mock function with given fields: id, name, notes, url, interval
func (_m *Client) RegisterCheck(id string, name string, notes string, url string, interval string) error {
	ret := _m.Called(id, name, notes, url, interval)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, string) error); ok {
		r0 = rf(id, name, notes, url, interval)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unregister provides a mock function with given fields:
func (_m *Client) Unregister() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
