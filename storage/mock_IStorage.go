// Code generated by mockery v2.43.2. DO NOT EDIT.

package storage

import mock "github.com/stretchr/testify/mock"

// MockStorage is an autogenerated mock type for the IStorage type
type MockStorage struct {
	mock.Mock
}

type MockStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStorage) EXPECT() *MockStorage_Expecter {
	return &MockStorage_Expecter{mock: &_m.Mock}
}

// EntityPath provides a mock function with given fields: collectionPath, entityRelativePath
func (_m *MockStorage) EntityPath(collectionPath string, entityRelativePath string) string {
	ret := _m.Called(collectionPath, entityRelativePath)

	if len(ret) == 0 {
		panic("no return value specified for EntityPath")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(collectionPath, entityRelativePath)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockStorage_EntityPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EntityPath'
type MockStorage_EntityPath_Call struct {
	*mock.Call
}

// EntityPath is a helper method to define mock.On call
//   - collectionPath string
//   - entityRelativePath string
func (_e *MockStorage_Expecter) EntityPath(collectionPath interface{}, entityRelativePath interface{}) *MockStorage_EntityPath_Call {
	return &MockStorage_EntityPath_Call{Call: _e.mock.On("EntityPath", collectionPath, entityRelativePath)}
}

func (_c *MockStorage_EntityPath_Call) Run(run func(collectionPath string, entityRelativePath string)) *MockStorage_EntityPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockStorage_EntityPath_Call) Return(_a0 string) *MockStorage_EntityPath_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorage_EntityPath_Call) RunAndReturn(run func(string, string) string) *MockStorage_EntityPath_Call {
	_c.Call.Return(run)
	return _c
}

// Main provides a mock function with given fields:
func (_m *MockStorage) Main() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Main")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorage_Main_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Main'
type MockStorage_Main_Call struct {
	*mock.Call
}

// Main is a helper method to define mock.On call
func (_e *MockStorage_Expecter) Main() *MockStorage_Main_Call {
	return &MockStorage_Main_Call{Call: _e.mock.On("Main")}
}

func (_c *MockStorage_Main_Call) Run(run func()) *MockStorage_Main_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStorage_Main_Call) Return(_a0 string, _a1 error) *MockStorage_Main_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorage_Main_Call) RunAndReturn(run func() (string, error)) *MockStorage_Main_Call {
	_c.Call.Return(run)
	return _c
}

// MakePaths provides a mock function with given fields: paths
func (_m *MockStorage) MakePaths(paths AbsPaths) error {
	ret := _m.Called(paths)

	if len(ret) == 0 {
		panic("no return value specified for MakePaths")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(AbsPaths) error); ok {
		r0 = rf(paths)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStorage_MakePaths_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MakePaths'
type MockStorage_MakePaths_Call struct {
	*mock.Call
}

// MakePaths is a helper method to define mock.On call
//   - paths AbsPaths
func (_e *MockStorage_Expecter) MakePaths(paths interface{}) *MockStorage_MakePaths_Call {
	return &MockStorage_MakePaths_Call{Call: _e.mock.On("MakePaths", paths)}
}

func (_c *MockStorage_MakePaths_Call) Run(run func(paths AbsPaths)) *MockStorage_MakePaths_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(AbsPaths))
	})
	return _c
}

func (_c *MockStorage_MakePaths_Call) Return(_a0 error) *MockStorage_MakePaths_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorage_MakePaths_Call) RunAndReturn(run func(AbsPaths) error) *MockStorage_MakePaths_Call {
	_c.Call.Return(run)
	return _c
}

// Paths provides a mock function with given fields: collectionName
func (_m *MockStorage) Paths(collectionName string) (*AbsPaths, error) {
	ret := _m.Called(collectionName)

	if len(ret) == 0 {
		panic("no return value specified for Paths")
	}

	var r0 *AbsPaths
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*AbsPaths, error)); ok {
		return rf(collectionName)
	}
	if rf, ok := ret.Get(0).(func(string) *AbsPaths); ok {
		r0 = rf(collectionName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AbsPaths)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(collectionName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorage_Paths_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Paths'
type MockStorage_Paths_Call struct {
	*mock.Call
}

// Paths is a helper method to define mock.On call
//   - collectionName string
func (_e *MockStorage_Expecter) Paths(collectionName interface{}) *MockStorage_Paths_Call {
	return &MockStorage_Paths_Call{Call: _e.mock.On("Paths", collectionName)}
}

func (_c *MockStorage_Paths_Call) Run(run func(collectionName string)) *MockStorage_Paths_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockStorage_Paths_Call) Return(_a0 *AbsPaths, _a1 error) *MockStorage_Paths_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorage_Paths_Call) RunAndReturn(run func(string) (*AbsPaths, error)) *MockStorage_Paths_Call {
	_c.Call.Return(run)
	return _c
}

// TmpMain provides a mock function with given fields:
func (_m *MockStorage) TmpMain() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TmpMain")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorage_TmpMain_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TmpMain'
type MockStorage_TmpMain_Call struct {
	*mock.Call
}

// TmpMain is a helper method to define mock.On call
func (_e *MockStorage_Expecter) TmpMain() *MockStorage_TmpMain_Call {
	return &MockStorage_TmpMain_Call{Call: _e.mock.On("TmpMain")}
}

func (_c *MockStorage_TmpMain_Call) Run(run func()) *MockStorage_TmpMain_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStorage_TmpMain_Call) Return(_a0 string, _a1 error) *MockStorage_TmpMain_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorage_TmpMain_Call) RunAndReturn(run func() (string, error)) *MockStorage_TmpMain_Call {
	_c.Call.Return(run)
	return _c
}

// TmpPaths provides a mock function with given fields: collectionName
func (_m *MockStorage) TmpPaths(collectionName string) (*AbsPaths, error) {
	ret := _m.Called(collectionName)

	if len(ret) == 0 {
		panic("no return value specified for TmpPaths")
	}

	var r0 *AbsPaths
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*AbsPaths, error)); ok {
		return rf(collectionName)
	}
	if rf, ok := ret.Get(0).(func(string) *AbsPaths); ok {
		r0 = rf(collectionName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AbsPaths)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(collectionName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorage_TmpPaths_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TmpPaths'
type MockStorage_TmpPaths_Call struct {
	*mock.Call
}

// TmpPaths is a helper method to define mock.On call
//   - collectionName string
func (_e *MockStorage_Expecter) TmpPaths(collectionName interface{}) *MockStorage_TmpPaths_Call {
	return &MockStorage_TmpPaths_Call{Call: _e.mock.On("TmpPaths", collectionName)}
}

func (_c *MockStorage_TmpPaths_Call) Run(run func(collectionName string)) *MockStorage_TmpPaths_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockStorage_TmpPaths_Call) Return(_a0 *AbsPaths, _a1 error) *MockStorage_TmpPaths_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorage_TmpPaths_Call) RunAndReturn(run func(string) (*AbsPaths, error)) *MockStorage_TmpPaths_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStorage creates a new instance of MockStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStorage {
	mock := &MockStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
