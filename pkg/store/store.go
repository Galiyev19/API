package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Open(data_base_url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", data_base_url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	s.db = db

	return db, nil
}

func (s *Store) Close() error {
	return nil
}
