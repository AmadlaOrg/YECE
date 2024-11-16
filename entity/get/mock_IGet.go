// Code generated by mockery v2.47.0. DO NOT EDIT.

package get

import (
	entity "github.com/AmadlaOrg/hery/entity"
	mock "github.com/stretchr/testify/mock"

	storage "github.com/AmadlaOrg/hery/storage"
)

// MockEntityGet is an autogenerated mock type for the IGet type
type MockEntityGet struct {
	mock.Mock
}

type MockEntityGet_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEntityGet) EXPECT() *MockEntityGet_Expecter {
	return &MockEntityGet_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: collectionName, storagePaths, entities
func (_m *MockEntityGet) Get(collectionName string, storagePaths *storage.AbsPaths, entities []string) error {
	ret := _m.Called(collectionName, storagePaths, entities)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *storage.AbsPaths, []string) error); ok {
		r0 = rf(collectionName, storagePaths, entities)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEntityGet_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockEntityGet_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - collectionName string
//   - storagePaths *storage.AbsPaths
//   - entities []string
func (_e *MockEntityGet_Expecter) Get(collectionName interface{}, storagePaths interface{}, entities interface{}) *MockEntityGet_Get_Call {
	return &MockEntityGet_Get_Call{Call: _e.mock.On("Get", collectionName, storagePaths, entities)}
}

func (_c *MockEntityGet_Get_Call) Run(run func(collectionName string, storagePaths *storage.AbsPaths, entities []string)) *MockEntityGet_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*storage.AbsPaths), args[2].([]string))
	})
	return _c
}

func (_c *MockEntityGet_Get_Call) Return(_a0 error) *MockEntityGet_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEntityGet_Get_Call) RunAndReturn(run func(string, *storage.AbsPaths, []string) error) *MockEntityGet_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetInTmp provides a mock function with given fields: collectionName, entities
func (_m *MockEntityGet) GetInTmp(collectionName string, entities []string) (storage.AbsPaths, error) {
	ret := _m.Called(collectionName, entities)

	if len(ret) == 0 {
		panic("no return value specified for GetInTmp")
	}

	var r0 storage.AbsPaths
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []string) (storage.AbsPaths, error)); ok {
		return rf(collectionName, entities)
	}
	if rf, ok := ret.Get(0).(func(string, []string) storage.AbsPaths); ok {
		r0 = rf(collectionName, entities)
	} else {
		r0 = ret.Get(0).(storage.AbsPaths)
	}

	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(collectionName, entities)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEntityGet_GetInTmp_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInTmp'
type MockEntityGet_GetInTmp_Call struct {
	*mock.Call
}

// GetInTmp is a helper method to define mock.On call
//   - collectionName string
//   - entities []string
func (_e *MockEntityGet_Expecter) GetInTmp(collectionName interface{}, entities interface{}) *MockEntityGet_GetInTmp_Call {
	return &MockEntityGet_GetInTmp_Call{Call: _e.mock.On("GetInTmp", collectionName, entities)}
}

func (_c *MockEntityGet_GetInTmp_Call) Run(run func(collectionName string, entities []string)) *MockEntityGet_GetInTmp_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]string))
	})
	return _c
}

func (_c *MockEntityGet_GetInTmp_Call) Return(_a0 storage.AbsPaths, _a1 error) *MockEntityGet_GetInTmp_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEntityGet_GetInTmp_Call) RunAndReturn(run func(string, []string) (storage.AbsPaths, error)) *MockEntityGet_GetInTmp_Call {
	_c.Call.Return(run)
	return _c
}

// download provides a mock function with given fields: collectionName, storagePaths, entitiesMeta
func (_m *MockEntityGet) download(collectionName string, storagePaths *storage.AbsPaths, entitiesMeta []entity.Entity) error {
	ret := _m.Called(collectionName, storagePaths, entitiesMeta)

	if len(ret) == 0 {
		panic("no return value specified for download")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *storage.AbsPaths, []entity.Entity) error); ok {
		r0 = rf(collectionName, storagePaths, entitiesMeta)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEntityGet_download_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'download'
type MockEntityGet_download_Call struct {
	*mock.Call
}

// download is a helper method to define mock.On call
//   - collectionName string
//   - storagePaths *storage.AbsPaths
//   - entitiesMeta []entity.Entity
func (_e *MockEntityGet_Expecter) download(collectionName interface{}, storagePaths interface{}, entitiesMeta interface{}) *MockEntityGet_download_Call {
	return &MockEntityGet_download_Call{Call: _e.mock.On("download", collectionName, storagePaths, entitiesMeta)}
}

func (_c *MockEntityGet_download_Call) Run(run func(collectionName string, storagePaths *storage.AbsPaths, entitiesMeta []entity.Entity)) *MockEntityGet_download_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*storage.AbsPaths), args[2].([]entity.Entity))
	})
	return _c
}

func (_c *MockEntityGet_download_Call) Return(_a0 error) *MockEntityGet_download_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEntityGet_download_Call) RunAndReturn(run func(string, *storage.AbsPaths, []entity.Entity) error) *MockEntityGet_download_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEntityGet creates a new instance of MockEntityGet. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEntityGet(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEntityGet {
	mock := &MockEntityGet{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
