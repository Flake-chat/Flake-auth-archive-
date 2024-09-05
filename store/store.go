package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	config  *Config
	db      *sql.DB
	userrep *Userrep
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DB)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Store) Close() error {
	s.db.Close()
	return nil
}

func (s *Store) User() *Userrep {
	if s.userrep != nil {
		return s.userrep
	}

	s.userrep = &Userrep{
		store: s,
	}

	return s.userrep
}
