// Code generated by mockery v2.47.0. DO NOT EDIT.

package validation

import (
	jsonschema "github.com/santhosh-tekuri/jsonschema/v6"
	mock "github.com/stretchr/testify/mock"
)

// MockEntityValidation is an autogenerated mock type for the IValidation type
type MockEntityValidation struct {
	mock.Mock
}

type MockEntityValidation_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEntityValidation) EXPECT() *MockEntityValidation_Expecter {
	return &MockEntityValidation_Expecter{mock: &_m.Mock}
}

// Entity provides a mock function with given fields: collectionName, schema, heryContent
func (_m *MockEntityValidation) Entity(collectionName string, schema *jsonschema.Schema, heryContent map[string]any) error {
	ret := _m.Called(collectionName, schema, heryContent)

	if len(ret) == 0 {
		panic("no return value specified for Entity")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *jsonschema.Schema, map[string]any) error); ok {
		r0 = rf(collectionName, schema, heryContent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEntityValidation_Entity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Entity'
type MockEntityValidation_Entity_Call struct {
	*mock.Call
}

// Entity is a helper method to define mock.On call
//   - collectionName string
//   - schema *jsonschema.Schema
//   - heryContent map[string]any
func (_e *MockEntityValidation_Expecter) Entity(collectionName interface{}, schema interface{}, heryContent interface{}) *MockEntityValidation_Entity_Call {
	return &MockEntityValidation_Entity_Call{Call: _e.mock.On("Entity", collectionName, schema, heryContent)}
}

func (_c *MockEntityValidation_Entity_Call) Run(run func(collectionName string, schema *jsonschema.Schema, heryContent map[string]any)) *MockEntityValidation_Entity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*jsonschema.Schema), args[2].(map[string]any))
	})
	return _c
}

func (_c *MockEntityValidation_Entity_Call) Return(_a0 error) *MockEntityValidation_Entity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEntityValidation_Entity_Call) RunAndReturn(run func(string, *jsonschema.Schema, map[string]any) error) *MockEntityValidation_Entity_Call {
	_c.Call.Return(run)
	return _c
}

// EntityUri provides a mock function with given fields: entityUrl
func (_m *MockEntityValidation) EntityUri(entityUrl string) bool {
	ret := _m.Called(entityUrl)

	if len(ret) == 0 {
		panic("no return value specified for EntityUri")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(entityUrl)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockEntityValidation_EntityUri_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EntityUri'
type MockEntityValidation_EntityUri_Call struct {
	*mock.Call
}

// EntityUri is a helper method to define mock.On call
//   - entityUrl string
func (_e *MockEntityValidation_Expecter) EntityUri(entityUrl interface{}) *MockEntityValidation_EntityUri_Call {
	return &MockEntityValidation_EntityUri_Call{Call: _e.mock.On("EntityUri", entityUrl)}
}

func (_c *MockEntityValidation_EntityUri_Call) Run(run func(entityUrl string)) *MockEntityValidation_EntityUri_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockEntityValidation_EntityUri_Call) Return(_a0 bool) *MockEntityValidation_EntityUri_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEntityValidation_EntityUri_Call) RunAndReturn(run func(string) bool) *MockEntityValidation_EntityUri_Call {
	_c.Call.Return(run)
	return _c
}

// RootEntity provides a mock function with given fields: rootSchema, selfSchema, heryContent
func (_m *MockEntityValidation) RootEntity(rootSchema *jsonschema.Schema, selfSchema *jsonschema.Schema, heryContent map[string]any) error {
	ret := _m.Called(rootSchema, selfSchema, heryContent)

	if len(ret) == 0 {
		panic("no return value specified for RootEntity")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*jsonschema.Schema, *jsonschema.Schema, map[string]any) error); ok {
		r0 = rf(rootSchema, selfSchema, heryContent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEntityValidation_RootEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RootEntity'
type MockEntityValidation_RootEntity_Call struct {
	*mock.Call
}

// RootEntity is a helper method to define mock.On call
//   - rootSchema *jsonschema.Schema
//   - selfSchema *jsonschema.Schema
//   - heryContent map[string]any
func (_e *MockEntityValidation_Expecter) RootEntity(rootSchema interface{}, selfSchema interface{}, heryContent interface{}) *MockEntityValidation_RootEntity_Call {
	return &MockEntityValidation_RootEntity_Call{Call: _e.mock.On("RootEntity", rootSchema, selfSchema, heryContent)}
}

func (_c *MockEntityValidation_RootEntity_Call) Run(run func(rootSchema *jsonschema.Schema, selfSchema *jsonschema.Schema, heryContent map[string]any)) *MockEntityValidation_RootEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*jsonschema.Schema), args[1].(*jsonschema.Schema), args[2].(map[string]any))
	})
	return _c
}

func (_c *MockEntityValidation_RootEntity_Call) Return(_a0 error) *MockEntityValidation_RootEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEntityValidation_RootEntity_Call) RunAndReturn(run func(*jsonschema.Schema, *jsonschema.Schema, map[string]any) error) *MockEntityValidation_RootEntity_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEntityValidation creates a new instance of MockEntityValidation. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEntityValidation(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEntityValidation {
	mock := &MockEntityValidation{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
