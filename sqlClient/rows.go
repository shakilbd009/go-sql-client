package sqlClient

import "database/sql"

type sqlRows struct {
	rows *sql.Rows
}
type rows interface {
	HasNext() bool
	Close() error
	Scan(dest ...interface{}) error
}

func (s *sqlRows) HasNext() bool {
	return s.rows.Next()
}

func (s *sqlRows) Close() error {
	return s.Close()
}

func (s *sqlRows) Scan(dest ...interface{}) error {
	return s.Scan(dest...)
}
