// Code generated by mockery v2.50.4. DO NOT EDIT.

package remote

import mock "github.com/stretchr/testify/mock"

// MockUtilGitRemote is an autogenerated mock type for the IRemote type
type MockUtilGitRemote struct {
	mock.Mock
}

type MockUtilGitRemote_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUtilGitRemote) EXPECT() *MockUtilGitRemote_Expecter {
	return &MockUtilGitRemote_Expecter{mock: &_m.Mock}
}

// CommitHeadHash provides a mock function with given fields: url
func (_m *MockUtilGitRemote) CommitHeadHash(url string) (string, error) {
	ret := _m.Called(url)

	if len(ret) == 0 {
		panic("no return value specified for CommitHeadHash")
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

// MockUtilGitRemote_CommitHeadHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CommitHeadHash'
type MockUtilGitRemote_CommitHeadHash_Call struct {
	*mock.Call
}

// CommitHeadHash is a helper method to define mock.On call
//   - url string
func (_e *MockUtilGitRemote_Expecter) CommitHeadHash(url interface{}) *MockUtilGitRemote_CommitHeadHash_Call {
	return &MockUtilGitRemote_CommitHeadHash_Call{Call: _e.mock.On("CommitHeadHash", url)}
}

func (_c *MockUtilGitRemote_CommitHeadHash_Call) Run(run func(url string)) *MockUtilGitRemote_CommitHeadHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockUtilGitRemote_CommitHeadHash_Call) Return(_a0 string, _a1 error) *MockUtilGitRemote_CommitHeadHash_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUtilGitRemote_CommitHeadHash_Call) RunAndReturn(run func(string) (string, error)) *MockUtilGitRemote_CommitHeadHash_Call {
	_c.Call.Return(run)
	return _c
}

// Tags provides a mock function with given fields: url
func (_m *MockUtilGitRemote) Tags(url string) ([]string, error) {
	ret := _m.Called(url)

	if len(ret) == 0 {
		panic("no return value specified for Tags")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]string, error)); ok {
		return rf(url)
	}
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUtilGitRemote_Tags_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Tags'
type MockUtilGitRemote_Tags_Call struct {
	*mock.Call
}

// Tags is a helper method to define mock.On call
//   - url string
func (_e *MockUtilGitRemote_Expecter) Tags(url interface{}) *MockUtilGitRemote_Tags_Call {
	return &MockUtilGitRemote_Tags_Call{Call: _e.mock.On("Tags", url)}
}

func (_c *MockUtilGitRemote_Tags_Call) Run(run func(url string)) *MockUtilGitRemote_Tags_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockUtilGitRemote_Tags_Call) Return(_a0 []string, _a1 error) *MockUtilGitRemote_Tags_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUtilGitRemote_Tags_Call) RunAndReturn(run func(string) ([]string, error)) *MockUtilGitRemote_Tags_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUtilGitRemote creates a new instance of MockUtilGitRemote. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUtilGitRemote(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUtilGitRemote {
	mock := &MockUtilGitRemote{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
