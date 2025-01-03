// Code generated by mockery v2.43.2. DO NOT EDIT.

package parser

import (
	database "github.com/AmadlaOrg/hery/cache/database"
	entity "github.com/AmadlaOrg/hery/entity"

	mock "github.com/stretchr/testify/mock"
)

// MockCacheParser is an autogenerated mock type for the IParser type
type MockCacheParser struct {
	mock.Mock
}

type MockCacheParser_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCacheParser) EXPECT() *MockCacheParser_Expecter {
	return &MockCacheParser_Expecter{mock: &_m.Mock}
}

// DatabaseRow provides a mock function with given fields: data
func (_m *MockCacheParser) DatabaseRow(data []byte) (entity.Entity, error) {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for DatabaseRow")
	}

	var r0 entity.Entity
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (entity.Entity, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func([]byte) entity.Entity); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(entity.Entity)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCacheParser_DatabaseRow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DatabaseRow'
type MockCacheParser_DatabaseRow_Call struct {
	*mock.Call
}

// DatabaseRow is a helper method to define mock.On call
//   - data []byte
func (_e *MockCacheParser_Expecter) DatabaseRow(data interface{}) *MockCacheParser_DatabaseRow_Call {
	return &MockCacheParser_DatabaseRow_Call{Call: _e.mock.On("DatabaseRow", data)}
}

func (_c *MockCacheParser_DatabaseRow_Call) Run(run func(data []byte)) *MockCacheParser_DatabaseRow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockCacheParser_DatabaseRow_Call) Return(_a0 entity.Entity, _a1 error) *MockCacheParser_DatabaseRow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCacheParser_DatabaseRow_Call) RunAndReturn(run func([]byte) (entity.Entity, error)) *MockCacheParser_DatabaseRow_Call {
	_c.Call.Return(run)
	return _c
}

// DatabaseTable provides a mock function with given fields: data
func (_m *MockCacheParser) DatabaseTable(data []byte) (entity.Entity, error) {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for DatabaseTable")
	}

	var r0 entity.Entity
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (entity.Entity, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func([]byte) entity.Entity); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(entity.Entity)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCacheParser_DatabaseTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DatabaseTable'
type MockCacheParser_DatabaseTable_Call struct {
	*mock.Call
}

// DatabaseTable is a helper method to define mock.On call
//   - data []byte
func (_e *MockCacheParser_Expecter) DatabaseTable(data interface{}) *MockCacheParser_DatabaseTable_Call {
	return &MockCacheParser_DatabaseTable_Call{Call: _e.mock.On("DatabaseTable", data)}
}

func (_c *MockCacheParser_DatabaseTable_Call) Run(run func(data []byte)) *MockCacheParser_DatabaseTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockCacheParser_DatabaseTable_Call) Return(_a0 entity.Entity, _a1 error) *MockCacheParser_DatabaseTable_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCacheParser_DatabaseTable_Call) RunAndReturn(run func([]byte) (entity.Entity, error)) *MockCacheParser_DatabaseTable_Call {
	_c.Call.Return(run)
	return _c
}

// Entity provides a mock function with given fields: _a0
func (_m *MockCacheParser) Entity(_a0 *entity.Entity) ([]database.Table, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Entity")
	}

	var r0 []database.Table
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Entity) ([]database.Table, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*entity.Entity) []database.Table); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.Table)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.Entity) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCacheParser_Entity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Entity'
type MockCacheParser_Entity_Call struct {
	*mock.Call
}

// Entity is a helper method to define mock.On call
//   - _a0 *entity.Entity
func (_e *MockCacheParser_Expecter) Entity(_a0 interface{}) *MockCacheParser_Entity_Call {
	return &MockCacheParser_Entity_Call{Call: _e.mock.On("Entity", _a0)}
}

func (_c *MockCacheParser_Entity_Call) Run(run func(_a0 *entity.Entity)) *MockCacheParser_Entity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Entity))
	})
	return _c
}

func (_c *MockCacheParser_Entity_Call) Return(_a0 []database.Table, _a1 error) *MockCacheParser_Entity_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCacheParser_Entity_Call) RunAndReturn(run func(*entity.Entity) ([]database.Table, error)) *MockCacheParser_Entity_Call {
	_c.Call.Return(run)
	return _c
}

// EntityToTableName provides a mock function with given fields: _a0
func (_m *MockCacheParser) EntityToTableName(_a0 string) string {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for EntityToTableName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockCacheParser_EntityToTableName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EntityToTableName'
type MockCacheParser_EntityToTableName_Call struct {
	*mock.Call
}

// EntityToTableName is a helper method to define mock.On call
//   - _a0 string
func (_e *MockCacheParser_Expecter) EntityToTableName(_a0 interface{}) *MockCacheParser_EntityToTableName_Call {
	return &MockCacheParser_EntityToTableName_Call{Call: _e.mock.On("EntityToTableName", _a0)}
}

func (_c *MockCacheParser_EntityToTableName_Call) Run(run func(_a0 string)) *MockCacheParser_EntityToTableName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockCacheParser_EntityToTableName_Call) Return(_a0 string) *MockCacheParser_EntityToTableName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCacheParser_EntityToTableName_Call) RunAndReturn(run func(string) string) *MockCacheParser_EntityToTableName_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCacheParser creates a new instance of MockCacheParser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCacheParser(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCacheParser {
	mock := &MockCacheParser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
