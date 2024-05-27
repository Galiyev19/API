package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateSqlDB(dbDriver, dbPath, migrationPath string) (*sql.DB, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open(dbDriver, dbPath)
	if err != nil {
		return nil, fmt.Errorf("OPEN DB: %w %s", err, op)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("PING ERROR: %w %s", err, op)
	}

	stmt, err := os.ReadFile(migrationPath)
	if err != nil {
		return nil, fmt.Errorf("READ FILE ERROR: %w %s", err, op)
	}

	_, err = db.Exec(string(stmt))
	if err != nil {
		return nil, fmt.Errorf("EXEC ERROR: %w %s", err, op)
	}
	return db, nil
}
