// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import parsers "github.com/testdouble/diplomat/parsers"

// ParagraphParser is an autogenerated mock type for the ParagraphParser type
type ParagraphParser struct {
	mock.Mock
}

// Parse provides a mock function with given fields: _a0
func (_m *ParagraphParser) Parse(_a0 []string) []parsers.Paragraph {
	ret := _m.Called(_a0)

	var r0 []parsers.Paragraph
	if rf, ok := ret.Get(0).(func([]string) []parsers.Paragraph); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]parsers.Paragraph)
		}
	}

	return r0
}

// ParseAll provides a mock function with given fields: _a0
func (_m *ParagraphParser) ParseAll(_a0 chan []string) chan parsers.Paragraph {
	ret := _m.Called(_a0)

	var r0 chan parsers.Paragraph
	if rf, ok := ret.Get(0).(func(chan []string) chan parsers.Paragraph); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan parsers.Paragraph)
		}
	}

	return r0
}
