// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

const Requester3_Get = "Get"

// Requester3 is an autogenerated mock type for the Requester3 type
type Requester3 struct {
	mock.Mock
}

// Get provides a mock function with given fields:
func (_m *Requester3) Get() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
