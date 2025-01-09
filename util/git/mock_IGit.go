// Code generated by mockery v2.50.4. DO NOT EDIT.

package git

import mock "github.com/stretchr/testify/mock"

// MockUtilGit is an autogenerated mock type for the IGit type
type MockUtilGit struct {
	mock.Mock
}

type MockUtilGit_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUtilGit) EXPECT() *MockUtilGit_Expecter {
	return &MockUtilGit_Expecter{mock: &_m.Mock}
}

// CheckoutTag provides a mock function with given fields: repoPath, tagName
func (_m *MockUtilGit) CheckoutTag(repoPath string, tagName string) error {
	ret := _m.Called(repoPath, tagName)

	if len(ret) == 0 {
		panic("no return value specified for CheckoutTag")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(repoPath, tagName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUtilGit_CheckoutTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckoutTag'
type MockUtilGit_CheckoutTag_Call struct {
	*mock.Call
}

// CheckoutTag is a helper method to define mock.On call
//   - repoPath string
//   - tagName string
func (_e *MockUtilGit_Expecter) CheckoutTag(repoPath interface{}, tagName interface{}) *MockUtilGit_CheckoutTag_Call {
	return &MockUtilGit_CheckoutTag_Call{Call: _e.mock.On("CheckoutTag", repoPath, tagName)}
}

func (_c *MockUtilGit_CheckoutTag_Call) Run(run func(repoPath string, tagName string)) *MockUtilGit_CheckoutTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockUtilGit_CheckoutTag_Call) Return(_a0 error) *MockUtilGit_CheckoutTag_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUtilGit_CheckoutTag_Call) RunAndReturn(run func(string, string) error) *MockUtilGit_CheckoutTag_Call {
	_c.Call.Return(run)
	return _c
}

// CommitHeadHash provides a mock function with given fields: repoPath
func (_m *MockUtilGit) CommitHeadHash(repoPath string) (string, error) {
	ret := _m.Called(repoPath)

	if len(ret) == 0 {
		panic("no return value specified for CommitHeadHash")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(repoPath)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(repoPath)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(repoPath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUtilGit_CommitHeadHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CommitHeadHash'
type MockUtilGit_CommitHeadHash_Call struct {
	*mock.Call
}

// CommitHeadHash is a helper method to define mock.On call
//   - repoPath string
func (_e *MockUtilGit_Expecter) CommitHeadHash(repoPath interface{}) *MockUtilGit_CommitHeadHash_Call {
	return &MockUtilGit_CommitHeadHash_Call{Call: _e.mock.On("CommitHeadHash", repoPath)}
}

func (_c *MockUtilGit_CommitHeadHash_Call) Run(run func(repoPath string)) *MockUtilGit_CommitHeadHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockUtilGit_CommitHeadHash_Call) Return(_a0 string, _a1 error) *MockUtilGit_CommitHeadHash_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUtilGit_CommitHeadHash_Call) RunAndReturn(run func(string) (string, error)) *MockUtilGit_CommitHeadHash_Call {
	_c.Call.Return(run)
	return _c
}

// FetchRepo provides a mock function with given fields: url, dest
func (_m *MockUtilGit) FetchRepo(url string, dest string) error {
	ret := _m.Called(url, dest)

	if len(ret) == 0 {
		panic("no return value specified for FetchRepo")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(url, dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUtilGit_FetchRepo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchRepo'
type MockUtilGit_FetchRepo_Call struct {
	*mock.Call
}

// FetchRepo is a helper method to define mock.On call
//   - url string
//   - dest string
func (_e *MockUtilGit_Expecter) FetchRepo(url interface{}, dest interface{}) *MockUtilGit_FetchRepo_Call {
	return &MockUtilGit_FetchRepo_Call{Call: _e.mock.On("FetchRepo", url, dest)}
}

func (_c *MockUtilGit_FetchRepo_Call) Run(run func(url string, dest string)) *MockUtilGit_FetchRepo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockUtilGit_FetchRepo_Call) Return(_a0 error) *MockUtilGit_FetchRepo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUtilGit_FetchRepo_Call) RunAndReturn(run func(string, string) error) *MockUtilGit_FetchRepo_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUtilGit creates a new instance of MockUtilGit. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUtilGit(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUtilGit {
	mock := &MockUtilGit{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
