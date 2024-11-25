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
	Update(table Table)
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

	s.queries.CreateTable = append(s.queries.CreateTable, fmt.Sprintf("%s\n%s", createTableSQL, sqlIndexes))
}

// Insert inserts records into the table
func (s *SDatabase) Insert(table Table) {
	var queries []string
	for _, row := range table.Rows {
		var (
			columnNames       []string
			valuesPlaceholder []string
			columnValues      []string
		)

		// Iterate over columns in a single row
		for rowColumnName, rowValue := range row {
			columnNames = append(columnNames, rowColumnName)
			valuesPlaceholder = append(valuesPlaceholder, "?")
			columnValues = append(columnValues, rowValue) // TODO: Maybe use a struct and then attach a function to parse
		}

		// Construct the query for this row
		query := fmt.Sprintf(
			`INSERT INTO %s (%s) VALUES (%s)`,
			table.Name,
			strings.Join(columnNames, ", "),
			strings.Join(valuesPlaceholder, ", "),
		)

		// Append the query to the list
		queries = append(queries, query)

		// Add to s.queries if needed
		queryInsert := QueryInsert{
			Query:  query,
			Values: columnValues,
		}
		s.queries.Insert = append(s.queries.Insert, queryInsert)
	}
}

// Update updates a record in the table
func (s *SDatabase) Update(table Table) {
	*s.queries = append(*s.queries, fmt.Sprintf(`UPDATE %s SET name = ? WHERE id = ?`, table.Name))
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
	*s.queries = append(*s.queries, fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, table.Name))
}

// DropTable drops the table from the database
func (s *SDatabase) DropTable(table Table) {
	*s.queries = append(*s.queries, fmt.Sprintf(`DROP TABLE IF EXISTS %s`, table.Name))
}

// Apply all the SQL scripts that are in a string array that are combined into one SQL script
func (s *SDatabase) Apply() error {
	if !s.IsInitialized() {
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
	}(stmt)

	return nil
}
