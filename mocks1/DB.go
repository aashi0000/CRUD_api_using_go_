// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

// AutoMigrate provides a mock function with given fields: dst
func (_m *DB) AutoMigrate(dst ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dst...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AutoMigrate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(dst...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: value
func (_m *DB) Create(value interface{}) *gorm.DB {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}) *gorm.DB); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: value, where
func (_m *DB) Delete(value interface{}, where ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, where...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(value, where...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Find provides a mock function with given fields: out, where
func (_m *DB) Find(out interface{}, where ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, out)
	_ca = append(_ca, where...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(out, where...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// First provides a mock function with given fields: out, where
func (_m *DB) First(out interface{}, where ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, out)
	_ca = append(_ca, where...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for First")
	}

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(out, where...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Model provides a mock function with given fields: value
func (_m *DB) Model(value interface{}) *gorm.DB {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for Model")
	}

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}) *gorm.DB); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Updates provides a mock function with given fields: value
func (_m *DB) Updates(value interface{}) *gorm.DB {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for Updates")
	}

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}) *gorm.DB); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Where provides a mock function with given fields: query, args
func (_m *DB) Where(query interface{}, args ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Where")
	}

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// NewDB creates a new instance of DB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *DB {
	mock := &DB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}