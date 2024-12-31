package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

// IDatabase defines the database interface
type IDatabase interface {
	Initialize() error
	Close() error
	IsInitialized() bool
	CreateTable(table Table)
	Insert(table Table)
	Update(table Table, where []map[string]any)
	Select(table Table, name string) (string, error)
	Delete(table Table, id int)
	DropTable(table Table)
	Apply() error
}

// SDatabase implements IDatabase
type SDatabase struct {
	queries *Queries
}

var (
	db          *sql.DB
	dbMutex     sync.Mutex
	initErr     error
	initialized bool
)

// Initialize establishes the database connection
func (s *SDatabase) Initialize() error {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	if db != nil {
		return nil // Already initialized successfully
	}

	dbPath := "/tmp/hery.test.cache"
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		initErr = fmt.Errorf("error opening database: %v", err)
		return initErr
	}

	// Set PRAGMA statements for performance
	//
	// Doc: https://stackoverflow.com/questions/57118674/go-sqlite3-with-journal-mode-wal-gives-database-is-locked-error
	_, err = db.Exec("PRAGMA journal_mode = WAL;")
	if err != nil {
		initErr = fmt.Errorf("error setting journal mode: %v", err)
		err := db.Close()
		if err != nil {
			return err
		}
		db = nil
		return initErr
	}

	// Configure the database connection pool
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(0)

	// Initialization successful
	initErr = nil
	initialized = true

	return nil
}

// Close closes the database connection
func (s *SDatabase) Close() error {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	if !initialized {
		// Database was never initialized; nothing to close
		return nil
	}

	if db != nil {
		err := db.Close()
		db = nil
		initialized = false
		return err
	}

	initialized = false
	return nil
}

// IsInitialized returns true if the database has been initialized
func (s *SDatabase) IsInitialized() bool {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	return initialized
}

func (s *SDatabase) AddQuery(action string, query Query) {

}

// CreateTable creates a new table
func (s *SDatabase) CreateTable(table Table) {
	var sqlColumns string
	for _, column := range table.Columns {
		var columnConstraints string
		for _, constraint := range column.Constraints {
			columnConstraints = constraint.ToSQL()
		}
		sqlColumn := fmt.Sprintf("\n%s %s %s,", column.ColumnName, column.DataType, columnConstraints)
		sqlColumns = fmt.Sprintf("%s %s", sqlColumns, sqlColumn)
	}

	sqlColumns = strings.TrimSuffix(sqlColumns, " ,")

	var sqlRelationships string
	for _, relationship := range table.Relationships {
		sqlRelationship := fmt.Sprintf(
			",\nFOREIGN KEY(%s) REFERENCES %s(%s)",
			relationship.ColumnName,
			relationship.ReferencesTableName,
			relationship.ReferencesColumnName)
		sqlRelationships = fmt.Sprintf("%s %s", sqlRelationships, sqlRelationship)
	}

	createTableSQL := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s%s\n);", table.Name, sqlColumns, sqlRelationships)

	var sqlIndexes string
	for _, index := range table.Columns {
		// TODO: `idx_` what is this? Is this needed?
		//createIndexSQL := fmt.Sprintf(`CREATE INDEX IF NOT EXISTS idx_%s_name ON %s(name);`, table.Name, table.Name)
		sqlIndexe := fmt.Sprintf(
			"CREATE INDEX IF NOT EXISTS idx_%s_%s ON %s(%s);",
			table.Name,
			index.ColumnName,
			table.Name,
			index.ColumnName)
		sqlIndexes = fmt.Sprintf("%s\n%s", sqlIndexes, sqlIndexe)
	}

	queryCreateTable := Query{
		Query: fmt.Sprintf("%s\n%s", createTableSQL, sqlIndexes),
	}

	s.queries.CreateTable = append(s.queries.CreateTable, queryCreateTable)
}

// Insert inserts records into the table
func (s *SDatabase) Insert(table Table) {
	var (
		columnNames       []string
		valuesPlaceholder []string
		columnValues      []string
	)

	// Iterate over columns in a single row
	/*for rowIndex, rowValue := range table.Rows {
		println(rowIndex)
		rowKey := rowValue["key"]
		println(rowKey)
		// TODO: Add validation to the name
		//columnNames = append(columnNames, rowColumnName)

		// For each loop it adds to a string list `?` for a safe query
		valuesPlaceholder = append(valuesPlaceholder, "?")
		//columnValues = append(columnValues, rowValue) // TODO: Maybe use a struct and then attach a function to parse
	}
	println(columnValues)*/

	// Iterate over rows in the table
	for rowIndex, row := range table.Rows {
		println("Row Index:", rowIndex)

		// Iterate over keys and values in the row
		for key, value := range row {
			println("Column Name:", key)
			println("Value:", value)

			// Add the column name to columnNames
			columnNames = append(columnNames, key)

			// Add a placeholder for each value
			valuesPlaceholder = append(valuesPlaceholder, "?")

			// Convert the value to a string and add it to columnValues
			columnValues = append(columnValues, fmt.Sprintf("%v", value))
		}
	}

	// Debug output
	println("Column Names:", strings.Join(columnNames, ", "))
	println("Placeholders:", strings.Join(valuesPlaceholder, ", "))
	println("Column Values:", strings.Join(columnValues, ", "))

	// Construct the query for this row
	query := fmt.Sprintf(
		`INSERT INTO %s (%s) VALUES (%s)`,
		table.Name,
		strings.Join(columnNames, ", "),
		strings.Join(valuesPlaceholder, ", "),
	)

	println(query)

	// Append the query to the list

	// TODO: Do I keep?
	//queries = append(queries, query)

	// Add to s.queries if needed
	/*queryInsert := Query{
		Query:  query,
		Values: columnValues,
	}
	s.queries.Insert = append(s.queries.Insert, queryInsert)*/
}

// Update updates a record in the table
func (s *SDatabase) Update(table Table, where []map[string]any) {
	/*var (
		columnsWhere       []string
		columnsWhereValues []string
		valuesPlaceholder  []string
		columnValues       []string
	)

	for _, whereColumnElement := range where {
		for columnName, value := range row {
			// TODO: Add validation to the name
			columnsWhere = append(columnsWhere, fmt.Sprintf("%s = ?", whereColumnName))
			columnsWhereValues = append(columnsWhereValues, valueToSQL(whereColumnValue))
		}
	}

	// Iterate over columns in a single row
	for rowColumnName, rowValue := range table.Rows {
		for columnName, value := range row {
			//columnNames = append(columnNames, rowColumnName)
			// TODO: Add validation to the name
			valuesPlaceholder = append(valuesPlaceholder, fmt.Sprintf("%s = ?,", rowColumnName))
			columnValues = append(columnValues, valueToSQL(rowValue)) // TODO: Maybe use a struct and then attach a function to parse
		}
	}

	values := make([]string, len(table.Rows)+len(columnsWhere))
	values = append(columnValues, columnsWhereValues...)
	queryUpdate := Query{
		Query:  fmt.Sprintf(`UPDATE %s SET name = ? WHERE id = ?`, table.Name),
		Values: values,
	}
	s.queries.Update = append(s.queries.Update, queryUpdate)*/
}

// Select retrieves a record from the table
func (s *SDatabase) Select(table Table, name string) (string, error) {
	if !s.IsInitialized() {
		return "", fmt.Errorf(ErrorDatabaseNotInitialized)
	}

	var result string
	querySQL := fmt.Sprintf(`SELECT name FROM %s WHERE name = ?`, table.Name)
	err := db.QueryRow(querySQL, name).Scan(&result)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil // No matching record found
		}
		return "", fmt.Errorf("error querying record: %v", err)
	}

	return result, nil
}

// Delete deletes a record from the table
func (s *SDatabase) Delete(table Table, id int) {
	//*s.queries = append(*s.queries, fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, table.Name))
}

// DropTable drops the table from the database
func (s *SDatabase) DropTable(table Table) {
	//*s.queries = append(*s.queries, fmt.Sprintf(`DROP TABLE IF EXISTS %s`, table.Name))
}

// Apply all the SQL scripts that are in a string array that are combined into one SQL script
func (s *SDatabase) Apply() error {
	/*if !s.IsInitialized() {
		return fmt.Errorf(ErrorDatabaseNotInitialized)
	}

	stmt, err := db.Prepare(mergeSqlQueries(s.queries))
	if err != nil {
		return fmt.Errorf("error preparing insert statement: %v", err)
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			// TODO:
			return
		}
	}(stmt)*/

	return nil
}
