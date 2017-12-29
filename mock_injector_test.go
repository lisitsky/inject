// Code generated by mockery v1.0.0
package inject

import mock "github.com/stretchr/testify/mock"
import reflect "reflect"

// MockInjector is an autogenerated mock type for the Injector type
type MockInjector struct {
	mock.Mock
}

// Apply provides a mock function with given fields: _a0
func (_m *MockInjector) Apply(_a0 interface{}) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Construct provides a mock function with given fields: _a0
func (_m *MockInjector) Construct(_a0 interface{}) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConstructLater provides a mock function with given fields: _a0
func (_m *MockInjector) ConstructLater(_a0 interface{}) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FinishConstruct provides a mock function with given fields:
func (_m *MockInjector) FinishConstruct() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: _a0
func (_m *MockInjector) Get(_a0 reflect.Type) reflect.Value {
	ret := _m.Called(_a0)

	var r0 reflect.Value
	if rf, ok := ret.Get(0).(func(reflect.Type) reflect.Value); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(reflect.Value)
	}

	return r0
}

// Invoke provides a mock function with given fields: _a0
func (_m *MockInjector) Invoke(_a0 interface{}) ([]reflect.Value, error) {
	ret := _m.Called(_a0)

	var r0 []reflect.Value
	if rf, ok := ret.Get(0).(func(interface{}) []reflect.Value); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reflect.Value)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Map provides a mock function with given fields: _a0
func (_m *MockInjector) Map(_a0 interface{}) TypeMapper {
	ret := _m.Called(_a0)

	var r0 TypeMapper
	if rf, ok := ret.Get(0).(func(interface{}) TypeMapper); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(TypeMapper)
		}
	}

	return r0
}

// MapTo provides a mock function with given fields: _a0, _a1
func (_m *MockInjector) MapTo(_a0 interface{}, _a1 interface{}) TypeMapper {
	ret := _m.Called(_a0, _a1)

	var r0 TypeMapper
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) TypeMapper); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(TypeMapper)
		}
	}

	return r0
}

// Provide provides a mock function with given fields: _a0
func (_m *MockInjector) Provide(_a0 interface{}) TypeMapper {
	ret := _m.Called(_a0)

	var r0 TypeMapper
	if rf, ok := ret.Get(0).(func(interface{}) TypeMapper); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(TypeMapper)
		}
	}

	return r0
}

// Set provides a mock function with given fields: _a0, _a1
func (_m *MockInjector) Set(_a0 reflect.Type, _a1 reflect.Value) TypeMapper {
	ret := _m.Called(_a0, _a1)

	var r0 TypeMapper
	if rf, ok := ret.Get(0).(func(reflect.Type, reflect.Value) TypeMapper); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(TypeMapper)
		}
	}

	return r0
}

// SetParent provides a mock function with given fields: _a0
func (_m *MockInjector) SetParent(_a0 Injector) {
	_m.Called(_a0)
}