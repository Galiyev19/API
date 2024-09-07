package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func InitDB(storeDriver, connStr, migrationPath string) (*sql.DB, error) {
	db, err := sql.Open(storeDriver, connStr)
	if err != nil {
		return nil, fmt.Errorf("open DB: %v", err)
	}
	defer db.Close()

	stmt, err := os.ReadFile(migrationPath)
	if err != nil {
		return nil, fmt.Errorf("read migration: %v", err)
	}

	_, err = db.Exec(string(stmt))
	if err != nil {
		return nil, fmt.Errorf("exec db: %v", err)
	}

	return db, nil
}
