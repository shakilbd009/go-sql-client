package sqlClient

import (
	"database/sql"
	"fmt"
	"os"
)

const (
	goEnv = "GO_ENV"
	prod  = "production"
)

var (
	isMock   bool
	dbClient SqlClient
)

type client struct {
	db *sql.DB
}

type SqlClient interface {
	Query(query string, args ...interface{}) (rows, error)
}

func isProduction() bool {
	return os.Getenv(goEnv) == prod
}

func StartMockupServer() {
	isMock = true
}

func StopMockupServer() {
	isMock = false
}

func Open(driverName, dataSourceName string) (SqlClient, error) {

	if isMock && !isProduction() {
		dbClient = &clientMock{}
		return dbClient, nil
	}
	if driverName == "" {
		return nil, fmt.Errorf("invalid driver name")
	}
	database, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	dbClient = &client{
		db: database,
	}
	return dbClient, nil
}

func (c *client) Query(query string, args ...interface{}) (rows, error) {
	rows, err := c.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	sqlrows := &sqlRows{
		rows: rows,
	}
	return sqlrows, nil
}
