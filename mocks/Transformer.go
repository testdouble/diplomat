// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import parsers "github.com/testdouble/diplomat/parsers"

// Transformer is an autogenerated mock type for the Transformer type
type Transformer struct {
	mock.Mock
}

// Transform provides a mock function with given fields: _a0
func (_m *Transformer) Transform(_a0 parsers.Test) (parsers.Test, error) {
	ret := _m.Called(_a0)

	var r0 parsers.Test
	if rf, ok := ret.Get(0).(func(parsers.Test) parsers.Test); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(parsers.Test)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(parsers.Test) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransformAll provides a mock function with given fields: _a0, _a1
func (_m *Transformer) TransformAll(_a0 chan parsers.Test, _a1 chan error) chan parsers.Test {
	ret := _m.Called(_a0, _a1)

	var r0 chan parsers.Test
	if rf, ok := ret.Get(0).(func(chan parsers.Test, chan error) chan parsers.Test); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan parsers.Test)
		}
	}

	return r0
}