// Code generated by mockery v2.43.2. DO NOT EDIT.

package schema

import (
	jsonschema "github.com/santhosh-tekuri/jsonschema/v6"
	mock "github.com/stretchr/testify/mock"
)

// MockEntitySchema is an autogenerated mock type for the ISchema type
type MockEntitySchema struct {
	mock.Mock
}

type MockEntitySchema_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEntitySchema) EXPECT() *MockEntitySchema_Expecter {
	return &MockEntitySchema_Expecter{mock: &_m.Mock}
}

// ExtractBody provides a mock function with given fields: heryContent
func (_m *MockEntitySchema) ExtractBody(heryContent map[string]interface{}) map[string]interface{} {
	ret := _m.Called(heryContent)

	if len(ret) == 0 {
		panic("no return value specified for ExtractBody")
	}

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(map[string]interface{}) map[string]interface{}); ok {
		r0 = rf(heryContent)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// MockEntitySchema_ExtractBody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExtractBody'
type MockEntitySchema_ExtractBody_Call struct {
	*mock.Call
}

// ExtractBody is a helper method to define mock.On call
//   - heryContent map[string]interface{}
func (_e *MockEntitySchema_Expecter) ExtractBody(heryContent interface{}) *MockEntitySchema_ExtractBody_Call {
	return &MockEntitySchema_ExtractBody_Call{Call: _e.mock.On("ExtractBody", heryContent)}
}

func (_c *MockEntitySchema_ExtractBody_Call) Run(run func(heryContent map[string]interface{})) *MockEntitySchema_ExtractBody_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]interface{}))
	})
	return _c
}

func (_c *MockEntitySchema_ExtractBody_Call) Return(_a0 map[string]interface{}) *MockEntitySchema_ExtractBody_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEntitySchema_ExtractBody_Call) RunAndReturn(run func(map[string]interface{}) map[string]interface{}) *MockEntitySchema_ExtractBody_Call {
	_c.Call.Return(run)
	return _c
}

// GenerateSchemaPath provides a mock function with given fields: collectionName, entityPath
func (_m *MockEntitySchema) GenerateSchemaPath(collectionName string, entityPath string) string {
	ret := _m.Called(collectionName, entityPath)

	if len(ret) == 0 {
		panic("no return value specified for GenerateSchemaPath")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(collectionName, entityPath)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockEntitySchema_GenerateSchemaPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateSchemaPath'
type MockEntitySchema_GenerateSchemaPath_Call struct {
	*mock.Call
}

// GenerateSchemaPath is a helper method to define mock.On call
//   - collectionName string
//   - entityPath string
func (_e *MockEntitySchema_Expecter) GenerateSchemaPath(collectionName interface{}, entityPath interface{}) *MockEntitySchema_GenerateSchemaPath_Call {
	return &MockEntitySchema_GenerateSchemaPath_Call{Call: _e.mock.On("GenerateSchemaPath", collectionName, entityPath)}
}

func (_c *MockEntitySchema_GenerateSchemaPath_Call) Run(run func(collectionName string, entityPath string)) *MockEntitySchema_GenerateSchemaPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockEntitySchema_GenerateSchemaPath_Call) Return(_a0 string) *MockEntitySchema_GenerateSchemaPath_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEntitySchema_GenerateSchemaPath_Call) RunAndReturn(run func(string, string) string) *MockEntitySchema_GenerateSchemaPath_Call {
	_c.Call.Return(run)
	return _c
}

// GenerateURN provides a mock function with given fields: urnPrefix, entityUri
func (_m *MockEntitySchema) GenerateURN(urnPrefix string, entityUri string) string {
	ret := _m.Called(urnPrefix, entityUri)

	if len(ret) == 0 {
		panic("no return value specified for GenerateURN")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(urnPrefix, entityUri)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockEntitySchema_GenerateURN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateURN'
type MockEntitySchema_GenerateURN_Call struct {
	*mock.Call
}

// GenerateURN is a helper method to define mock.On call
//   - urnPrefix string
//   - entityUri string
func (_e *MockEntitySchema_Expecter) GenerateURN(urnPrefix interface{}, entityUri interface{}) *MockEntitySchema_GenerateURN_Call {
	return &MockEntitySchema_GenerateURN_Call{Call: _e.mock.On("GenerateURN", urnPrefix, entityUri)}
}

func (_c *MockEntitySchema_GenerateURN_Call) Run(run func(urnPrefix string, entityUri string)) *MockEntitySchema_GenerateURN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockEntitySchema_GenerateURN_Call) Return(_a0 string) *MockEntitySchema_GenerateURN_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEntitySchema_GenerateURN_Call) RunAndReturn(run func(string, string) string) *MockEntitySchema_GenerateURN_Call {
	_c.Call.Return(run)
	return _c
}

// GenerateURNPrefix provides a mock function with given fields: collectionName
func (_m *MockEntitySchema) GenerateURNPrefix(collectionName string) string {
	ret := _m.Called(collectionName)

	if len(ret) == 0 {
		panic("no return value specified for GenerateURNPrefix")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(collectionName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockEntitySchema_GenerateURNPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateURNPrefix'
type MockEntitySchema_GenerateURNPrefix_Call struct {
	*mock.Call
}

// GenerateURNPrefix is a helper method to define mock.On call
//   - collectionName string
func (_e *MockEntitySchema_Expecter) GenerateURNPrefix(collectionName interface{}) *MockEntitySchema_GenerateURNPrefix_Call {
	return &MockEntitySchema_GenerateURNPrefix_Call{Call: _e.mock.On("GenerateURNPrefix", collectionName)}
}

func (_c *MockEntitySchema_GenerateURNPrefix_Call) Run(run func(collectionName string)) *MockEntitySchema_GenerateURNPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockEntitySchema_GenerateURNPrefix_Call) Return(_a0 string) *MockEntitySchema_GenerateURNPrefix_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEntitySchema_GenerateURNPrefix_Call) RunAndReturn(run func(string) string) *MockEntitySchema_GenerateURNPrefix_Call {
	_c.Call.Return(run)
	return _c
}

// Load provides a mock function with given fields: schemaPath
func (_m *MockEntitySchema) Load(schemaPath string) (*jsonschema.Schema, error) {
	ret := _m.Called(schemaPath)

	if len(ret) == 0 {
		panic("no return value specified for Load")
	}

	var r0 *jsonschema.Schema
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*jsonschema.Schema, error)); ok {
		return rf(schemaPath)
	}
	if rf, ok := ret.Get(0).(func(string) *jsonschema.Schema); ok {
		r0 = rf(schemaPath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jsonschema.Schema)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(schemaPath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEntitySchema_Load_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Load'
type MockEntitySchema_Load_Call struct {
	*mock.Call
}

// Load is a helper method to define mock.On call
//   - schemaPath string
func (_e *MockEntitySchema_Expecter) Load(schemaPath interface{}) *MockEntitySchema_Load_Call {
	return &MockEntitySchema_Load_Call{Call: _e.mock.On("Load", schemaPath)}
}

func (_c *MockEntitySchema_Load_Call) Run(run func(schemaPath string)) *MockEntitySchema_Load_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockEntitySchema_Load_Call) Return(_a0 *jsonschema.Schema, _a1 error) *MockEntitySchema_Load_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEntitySchema_Load_Call) RunAndReturn(run func(string) (*jsonschema.Schema, error)) *MockEntitySchema_Load_Call {
	_c.Call.Return(run)
	return _c
}

// loadSchemaFile provides a mock function with given fields: schemaPath
func (_m *MockEntitySchema) loadSchemaFile(schemaPath string) (map[string]interface{}, error) {
	ret := _m.Called(schemaPath)

	if len(ret) == 0 {
		panic("no return value specified for loadSchemaFile")
	}

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (map[string]interface{}, error)); ok {
		return rf(schemaPath)
	}
	if rf, ok := ret.Get(0).(func(string) map[string]interface{}); ok {
		r0 = rf(schemaPath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(schemaPath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEntitySchema_loadSchemaFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'loadSchemaFile'
type MockEntitySchema_loadSchemaFile_Call struct {
	*mock.Call
}

// loadSchemaFile is a helper method to define mock.On call
//   - schemaPath string
func (_e *MockEntitySchema_Expecter) loadSchemaFile(schemaPath interface{}) *MockEntitySchema_loadSchemaFile_Call {
	return &MockEntitySchema_loadSchemaFile_Call{Call: _e.mock.On("loadSchemaFile", schemaPath)}
}

func (_c *MockEntitySchema_loadSchemaFile_Call) Run(run func(schemaPath string)) *MockEntitySchema_loadSchemaFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockEntitySchema_loadSchemaFile_Call) Return(_a0 map[string]interface{}, _a1 error) *MockEntitySchema_loadSchemaFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEntitySchema_loadSchemaFile_Call) RunAndReturn(run func(string) (map[string]interface{}, error)) *MockEntitySchema_loadSchemaFile_Call {
	_c.Call.Return(run)
	return _c
}

// mergeSchemas provides a mock function with given fields: baseSchema, mainSchema
func (_m *MockEntitySchema) mergeSchemas(baseSchema map[string]interface{}, mainSchema map[string]interface{}) map[string]interface{} {
	ret := _m.Called(baseSchema, mainSchema)

	if len(ret) == 0 {
		panic("no return value specified for mergeSchemas")
	}

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(map[string]interface{}, map[string]interface{}) map[string]interface{}); ok {
		r0 = rf(baseSchema, mainSchema)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// MockEntitySchema_mergeSchemas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mergeSchemas'
type MockEntitySchema_mergeSchemas_Call struct {
	*mock.Call
}

// mergeSchemas is a helper method to define mock.On call
//   - baseSchema map[string]interface{}
//   - mainSchema map[string]interface{}
func (_e *MockEntitySchema_Expecter) mergeSchemas(baseSchema interface{}, mainSchema interface{}) *MockEntitySchema_mergeSchemas_Call {
	return &MockEntitySchema_mergeSchemas_Call{Call: _e.mock.On("mergeSchemas", baseSchema, mainSchema)}
}

func (_c *MockEntitySchema_mergeSchemas_Call) Run(run func(baseSchema map[string]interface{}, mainSchema map[string]interface{})) *MockEntitySchema_mergeSchemas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]interface{}), args[1].(map[string]interface{}))
	})
	return _c
}

func (_c *MockEntitySchema_mergeSchemas_Call) Return(_a0 map[string]interface{}) *MockEntitySchema_mergeSchemas_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEntitySchema_mergeSchemas_Call) RunAndReturn(run func(map[string]interface{}, map[string]interface{}) map[string]interface{}) *MockEntitySchema_mergeSchemas_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEntitySchema creates a new instance of MockEntitySchema. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEntitySchema(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEntitySchema {
	mock := &MockEntitySchema{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
