package database

// Errors
const (
	ErrorClosingDatabase        = "closing database"
	ErrorDatabaseNotInitialized = "database not initialized"
)

type DataType string

// Data types
/*
SQLite 3 data type:
	Declared Type	       Type Affinity	Storage Class
	INTEGER	               INTEGER	        INTEGER
	TINYINT, SMALLINT	   INTEGER	        INTEGER
	MEDIUMINT, BIGINT	   INTEGER	        INTEGER
	INT	                   INTEGER	        INTEGER
	REAL, DOUBLE, FLOAT	   REAL	            REAL
	NUMERIC, DECIMAL	   NUMERIC	REAL or INTEGER (if possible)
	TEXT	               TEXT	            TEXT
	CHARACTER, VARCHAR	   TEXT	            TEXT
	CLOB	               TEXT	            TEXT
	BLOB	               BLOB	            BLOB
	BOOLEAN	               NUMERIC	         INTEGER (1 for true, 0 for false)
	DATE, DATETIME	       NUMERIC	TEXT, REAL, or INTEGER depending on the format
*/
const (
	DataTypeInteger    DataType = "INTEGER"
	DataTypeTinyint    DataType = "TINYINT"
	DataTypeBigInteger DataType = "BIGINT"
	DataTypeReal       DataType = "REAL"
	DataTypeNumeric    DataType = "NUMERIC"
	DataTypeDecimal    DataType = "DECIMAL"
	DataTypeBoolean    DataType = "BOOLEAN"
	DataTypeText       DataType = "TEXT"
	DataTypeCharacter  DataType = "CHARACTER"
	DataTypeVarchar    DataType = "VARCHAR"
	DataTypeClob       DataType = "CLOB"
	DataTypeBlob       DataType = "BLOB"
	DataTypeDate       DataType = "DATE"
	DataTypeDateTime   DataType = "DATETIME"
)

// Table is a basic representation in a struct of a table in a SQL DB
type Table struct {
	Name          string
	Columns       []Column
	Relationships []Relationships
	Rows          []map[string]any
}

// Column is a basic representation in a struct of a column in a SQL DB
type Column struct {
	ColumnName string
	DataType   DataType
	Constraint string
	Default    string
}

// Relationships so to create relationships
type Relationships struct {
	ColumnName           string
	ReferencesTableName  string
	ReferencesColumnName string
}
