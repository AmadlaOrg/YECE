// Code generated by mockery v2.50.4. DO NOT EDIT.

package database

import (
	sql "database/sql"

	mock "github.com/stretchr/testify/mock"
)

// MockSqlTx is an autogenerated mock type for the ISqlTx type
type MockSqlTx struct {
	mock.Mock
}

type MockSqlTx_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSqlTx) EXPECT() *MockSqlTx_Expecter {
	return &MockSqlTx_Expecter{mock: &_m.Mock}
}

// Commit provides a mock function with no fields
func (_m *MockSqlTx) Commit() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Commit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSqlTx_Commit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Commit'
type MockSqlTx_Commit_Call struct {
	*mock.Call
}

// Commit is a helper method to define mock.On call
func (_e *MockSqlTx_Expecter) Commit() *MockSqlTx_Commit_Call {
	return &MockSqlTx_Commit_Call{Call: _e.mock.On("Commit")}
}

func (_c *MockSqlTx_Commit_Call) Run(run func()) *MockSqlTx_Commit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSqlTx_Commit_Call) Return(_a0 error) *MockSqlTx_Commit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSqlTx_Commit_Call) RunAndReturn(run func() error) *MockSqlTx_Commit_Call {
	_c.Call.Return(run)
	return _c
}

// Exec provides a mock function with given fields: query, args
func (_m *MockSqlTx) Exec(query string, args ...any) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 sql.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...any) (sql.Result, error)); ok {
		return rf(query, args...)
	}
	if rf, ok := ret.Get(0).(func(string, ...any) sql.Result); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...any) error); ok {
		r1 = rf(query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSqlTx_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockSqlTx_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - query string
//   - args ...any
func (_e *MockSqlTx_Expecter) Exec(query interface{}, args ...interface{}) *MockSqlTx_Exec_Call {
	return &MockSqlTx_Exec_Call{Call: _e.mock.On("Exec",
		append([]interface{}{query}, args...)...)}
}

func (_c *MockSqlTx_Exec_Call) Run(run func(query string, args ...any)) *MockSqlTx_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]any, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(any)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSqlTx_Exec_Call) Return(_a0 sql.Result, _a1 error) *MockSqlTx_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSqlTx_Exec_Call) RunAndReturn(run func(string, ...any) (sql.Result, error)) *MockSqlTx_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// Rollback provides a mock function with no fields
func (_m *MockSqlTx) Rollback() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Rollback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSqlTx_Rollback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rollback'
type MockSqlTx_Rollback_Call struct {
	*mock.Call
}

// Rollback is a helper method to define mock.On call
func (_e *MockSqlTx_Expecter) Rollback() *MockSqlTx_Rollback_Call {
	return &MockSqlTx_Rollback_Call{Call: _e.mock.On("Rollback")}
}

func (_c *MockSqlTx_Rollback_Call) Run(run func()) *MockSqlTx_Rollback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSqlTx_Rollback_Call) Return(_a0 error) *MockSqlTx_Rollback_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSqlTx_Rollback_Call) RunAndReturn(run func() error) *MockSqlTx_Rollback_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSqlTx creates a new instance of MockSqlTx. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSqlTx(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSqlTx {
	mock := &MockSqlTx{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
