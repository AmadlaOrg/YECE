// Code generated by mockery v2.50.4. DO NOT EDIT.

package version

import mock "github.com/stretchr/testify/mock"

// MockEntityVersion is an autogenerated mock type for the IVersion type
type MockEntityVersion struct {
	mock.Mock
}

type MockEntityVersion_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEntityVersion) EXPECT() *MockEntityVersion_Expecter {
	return &MockEntityVersion_Expecter{mock: &_m.Mock}
}

// Extract provides a mock function with given fields: url
func (_m *MockEntityVersion) Extract(url string) (string, error) {
	ret := _m.Called(url)

	if len(ret) == 0 {
		panic("no return value specified for Extract")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(url)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEntityVersion_Extract_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Extract'
type MockEntityVersion_Extract_Call struct {
	*mock.Call
}

// Extract is a helper method to define mock.On call
//   - url string
func (_e *MockEntityVersion_Expecter) Extract(url interface{}) *MockEntityVersion_Extract_Call {
	return &MockEntityVersion_Extract_Call{Call: _e.mock.On("Extract", url)}
}

func (_c *MockEntityVersion_Extract_Call) Run(run func(url string)) *MockEntityVersion_Extract_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockEntityVersion_Extract_Call) Return(_a0 string, _a1 error) *MockEntityVersion_Extract_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEntityVersion_Extract_Call) RunAndReturn(run func(string) (string, error)) *MockEntityVersion_Extract_Call {
	_c.Call.Return(run)
	return _c
}

// GeneratePseudo provides a mock function with given fields: entityFullRepoUrl
func (_m *MockEntityVersion) GeneratePseudo(entityFullRepoUrl string) (string, error) {
	ret := _m.Called(entityFullRepoUrl)

	if len(ret) == 0 {
		panic("no return value specified for GeneratePseudo")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(entityFullRepoUrl)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(entityFullRepoUrl)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(entityFullRepoUrl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEntityVersion_GeneratePseudo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GeneratePseudo'
type MockEntityVersion_GeneratePseudo_Call struct {
	*mock.Call
}

// GeneratePseudo is a helper method to define mock.On call
//   - entityFullRepoUrl string
func (_e *MockEntityVersion_Expecter) GeneratePseudo(entityFullRepoUrl interface{}) *MockEntityVersion_GeneratePseudo_Call {
	return &MockEntityVersion_GeneratePseudo_Call{Call: _e.mock.On("GeneratePseudo", entityFullRepoUrl)}
}

func (_c *MockEntityVersion_GeneratePseudo_Call) Run(run func(entityFullRepoUrl string)) *MockEntityVersion_GeneratePseudo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockEntityVersion_GeneratePseudo_Call) Return(_a0 string, _a1 error) *MockEntityVersion_GeneratePseudo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEntityVersion_GeneratePseudo_Call) RunAndReturn(run func(string) (string, error)) *MockEntityVersion_GeneratePseudo_Call {
	_c.Call.Return(run)
	return _c
}

// Latest provides a mock function with given fields: versions
func (_m *MockEntityVersion) Latest(versions []string) (string, error) {
	ret := _m.Called(versions)

	if len(ret) == 0 {
		panic("no return value specified for Latest")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) (string, error)); ok {
		return rf(versions)
	}
	if rf, ok := ret.Get(0).(func([]string) string); ok {
		r0 = rf(versions)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(versions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEntityVersion_Latest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Latest'
type MockEntityVersion_Latest_Call struct {
	*mock.Call
}

// Latest is a helper method to define mock.On call
//   - versions []string
func (_e *MockEntityVersion_Expecter) Latest(versions interface{}) *MockEntityVersion_Latest_Call {
	return &MockEntityVersion_Latest_Call{Call: _e.mock.On("Latest", versions)}
}

func (_c *MockEntityVersion_Latest_Call) Run(run func(versions []string)) *MockEntityVersion_Latest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *MockEntityVersion_Latest_Call) Return(_a0 string, _a1 error) *MockEntityVersion_Latest_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEntityVersion_Latest_Call) RunAndReturn(run func([]string) (string, error)) *MockEntityVersion_Latest_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: entityUrlPath
func (_m *MockEntityVersion) List(entityUrlPath string) ([]string, error) {
	ret := _m.Called(entityUrlPath)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]string, error)); ok {
		return rf(entityUrlPath)
	}
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(entityUrlPath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(entityUrlPath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEntityVersion_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockEntityVersion_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - entityUrlPath string
func (_e *MockEntityVersion_Expecter) List(entityUrlPath interface{}) *MockEntityVersion_List_Call {
	return &MockEntityVersion_List_Call{Call: _e.mock.On("List", entityUrlPath)}
}

func (_c *MockEntityVersion_List_Call) Run(run func(entityUrlPath string)) *MockEntityVersion_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockEntityVersion_List_Call) Return(_a0 []string, _a1 error) *MockEntityVersion_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEntityVersion_List_Call) RunAndReturn(run func(string) ([]string, error)) *MockEntityVersion_List_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEntityVersion creates a new instance of MockEntityVersion. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEntityVersion(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEntityVersion {
	mock := &MockEntityVersion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
