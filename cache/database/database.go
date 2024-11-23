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
	CreateTable(table Table) error
	Insert(table Table, names []string) error
	Update(table Table, id int, newName string) error
	Select(table Table, name string) (string, error)
	Delete(table Table, id int) error
	DropTable(table Table) error
}

// SDatabase implements IDatabase
type SDatabase struct{}

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

	dbPath := "/tmp/amadla.cache"
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
func (s *SDatabase) CreateTable(table Table) string {
	var sqlColumns string
	for _, column := range table.Columns {
		sqlColumn := fmt.Sprintf(",\n%s %s %s", column.ColumnName, column.DataType, column.Constraint)
		sqlColumns = fmt.Sprintf("%s %s", sqlColumns, sqlColumn)
	}

	sqlColumns = strings.TrimPrefix(sqlColumns, ",")

	var sqlRelationships string
	for _, relationship := range table.Relationships {
		sqlRelationship := fmt.Sprintf(
			",\nFOREIGN KEY(%s) REFERENCES %s(%s)",
			relationship.ColumnName,
			relationship.ReferencesTableName,
			relationship.ReferencesColumnName)
		sqlRelationships = fmt.Sprintf("%s %s", sqlRelationships, sqlRelationship)
	}

	createTableSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s%s);`, table.Name, sqlColumns, sqlRelationships)

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

	return fmt.Sprintf("%s\n%s", createTableSQL, sqlIndexes)
}

// Insert inserts records into the table
func (s *SDatabase) Insert(table Table, names []string) string {
	return fmt.Sprintf(`INSERT INTO %s (name) VALUES (?)`, table.Name)
}

// Update updates a record in the table
func (s *SDatabase) Update(table Table, id int, newName string) string {
	return fmt.Sprintf(`UPDATE %s SET name = ? WHERE id = ?`, table.Name)
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
func (s *SDatabase) Delete(table Table, id int) string {
	return fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, table.Name)
}

// DropTable drops the table from the database
func (s *SDatabase) DropTable(table Table) string {
	return fmt.Sprintf(`DROP TABLE IF EXISTS %s`, table.Name)
}

// Apply all the SQL scripts that are in a string array that are combined into one SQL script
func (s *SDatabase) Apply(sqlQueries *[]string) error {
	if !s.IsInitialized() {
		return fmt.Errorf(ErrorDatabaseNotInitialized)
	}

	stmt, err := db.Prepare(mergeSqlQueries(sqlQueries))
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
