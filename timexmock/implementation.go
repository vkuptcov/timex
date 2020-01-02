// Code generated by mockery v1.0.0. DO NOT EDIT.

package timexmock

import (
	time "time"

	mock "github.com/stretchr/testify/mock"

	timex "github.com/cabify/timex"
)

// Implementation is an autogenerated mock type for the Implementation type
type Implementation struct {
	mock.Mock
}

// After provides a mock function with given fields: d
func (_m *Implementation) After(d time.Duration) <-chan time.Time {
	ret := _m.Called(d)

	var r0 <-chan time.Time
	if rf, ok := ret.Get(0).(func(time.Duration) <-chan time.Time); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan time.Time)
		}
	}

	return r0
}

// AfterFunc provides a mock function with given fields: d, f
func (_m *Implementation) AfterFunc(d time.Duration, f func()) timex.Timer {
	ret := _m.Called(d, f)

	var r0 timex.Timer
	if rf, ok := ret.Get(0).(func(time.Duration, func()) timex.Timer); ok {
		r0 = rf(d, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(timex.Timer)
		}
	}

	return r0
}

// NewTicker provides a mock function with given fields: d
func (_m *Implementation) NewTicker(d time.Duration) timex.Ticker {
	ret := _m.Called(d)

	var r0 timex.Ticker
	if rf, ok := ret.Get(0).(func(time.Duration) timex.Ticker); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(timex.Ticker)
		}
	}

	return r0
}

// NewTimer provides a mock function with given fields: d
func (_m *Implementation) NewTimer(d time.Duration) timex.Timer {
	ret := _m.Called(d)

	var r0 timex.Timer
	if rf, ok := ret.Get(0).(func(time.Duration) timex.Timer); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(timex.Timer)
		}
	}

	return r0
}

// Now provides a mock function with given fields:
func (_m *Implementation) Now() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// Since provides a mock function with given fields: t
func (_m *Implementation) Since(t time.Time) time.Duration {
	ret := _m.Called(t)

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func(time.Time) time.Duration); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// Sleep provides a mock function with given fields: d
func (_m *Implementation) Sleep(d time.Duration) {
	_m.Called(d)
}

// Until provides a mock function with given fields: t
func (_m *Implementation) Until(t time.Time) time.Duration {
	ret := _m.Called(t)

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func(time.Time) time.Duration); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}