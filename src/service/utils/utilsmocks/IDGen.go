// Code generated by mockery v1.0.0. DO NOT EDIT.

package utilsmocks

import mock "github.com/stretchr/testify/mock"

// IDGen is an autogenerated mock type for the IDGen type
type IDGen struct {
	mock.Mock
}

// NewID provides a mock function with given fields:
func (_m *IDGen) NewID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
