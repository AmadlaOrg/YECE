// Code generated by mockery v2.50.3. DO NOT EDIT.

package database

import mock "github.com/stretchr/testify/mock"

// MockSyncLocker is an autogenerated mock type for the Locker type
type MockSyncLocker struct {
	mock.Mock
}

type MockSyncLocker_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSyncLocker) EXPECT() *MockSyncLocker_Expecter {
	return &MockSyncLocker_Expecter{mock: &_m.Mock}
}

// Lock provides a mock function with no fields
func (_m *MockSyncLocker) Lock() {
	_m.Called()
}

// MockSyncLocker_Lock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Lock'
type MockSyncLocker_Lock_Call struct {
	*mock.Call
}

// Lock is a helper method to define mock.On call
func (_e *MockSyncLocker_Expecter) Lock() *MockSyncLocker_Lock_Call {
	return &MockSyncLocker_Lock_Call{Call: _e.mock.On("Lock")}
}

func (_c *MockSyncLocker_Lock_Call) Run(run func()) *MockSyncLocker_Lock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSyncLocker_Lock_Call) Return() *MockSyncLocker_Lock_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSyncLocker_Lock_Call) RunAndReturn(run func()) *MockSyncLocker_Lock_Call {
	_c.Run(run)
	return _c
}

// Unlock provides a mock function with no fields
func (_m *MockSyncLocker) Unlock() {
	_m.Called()
}

// MockSyncLocker_Unlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unlock'
type MockSyncLocker_Unlock_Call struct {
	*mock.Call
}

// Unlock is a helper method to define mock.On call
func (_e *MockSyncLocker_Expecter) Unlock() *MockSyncLocker_Unlock_Call {
	return &MockSyncLocker_Unlock_Call{Call: _e.mock.On("Unlock")}
}

func (_c *MockSyncLocker_Unlock_Call) Run(run func()) *MockSyncLocker_Unlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSyncLocker_Unlock_Call) Return() *MockSyncLocker_Unlock_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSyncLocker_Unlock_Call) RunAndReturn(run func()) *MockSyncLocker_Unlock_Call {
	_c.Run(run)
	return _c
}

// NewMockSyncLocker creates a new instance of MockSyncLocker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSyncLocker(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSyncLocker {
	mock := &MockSyncLocker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
