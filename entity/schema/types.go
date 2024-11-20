package schema

import "github.com/santhosh-tekuri/jsonschema/v6"

const (
	EntityJsonSchemaFileName = "schema.hery.json"
	EntityJsonSchemaIdURN    = `^urn:([a-z0-9][a-z0-9-]{0,31}):([a-z0-9][a-z0-9-]+):([a-zA-Z0-9\-.:]+):([a-zA-Z0-9\-.]+)$`
)

// JSON-Schema data type
/*
Convert JSON-Schema Types to SQLite 3 data types:
	string.
	number.
	integer.
	object.
	array.
	boolean.
	null.
*/
const (
	DataTypeString  = "string"
	DataTypeNumber  = "number"
	DataTypeInteger = "integer"
	DataTypeObject  = "object"
	DataTypeArray   = "array"
	DataTypeBoolean = "boolean"
	DateTypeNull    = "null"
)

// Schema different data
type Schema struct {
	CompiledSchema *jsonschema.Schema
	SchemaPath     string
	SchemaName     string
	SchemaId       string
	Schema         map[string]any
}
