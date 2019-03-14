// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import parsers "github.com/testdouble/diplomat/parsers"

// SpecParser is an autogenerated mock type for the SpecParser type
type SpecParser struct {
	mock.Mock
}

// Parse provides a mock function with given fields: _a0, _a1
func (_m *SpecParser) Parse(_a0 chan string, _a1 chan error) chan parsers.Spec {
	ret := _m.Called(_a0, _a1)

	var r0 chan parsers.Spec
	if rf, ok := ret.Get(0).(func(chan string, chan error) chan parsers.Spec); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan parsers.Spec)
		}
	}

	return r0
}
