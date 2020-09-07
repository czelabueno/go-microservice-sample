package database

import (
	"database/sql"

	"go-restapi-mysql/internal/logs"
)

// MySQLClient model struct to instance
type MySQLClient struct {
	*sql.DB
}

// NewMySQLClient function return database tenant connection
func NewMySQLClient(source string) *sql.DB {
	db, err := sql.Open("mysql", source)
	if err != nil {
		logs.Log().Errorf("Cant create DB tenart: %s", err.Error()) // '_=' for ignore deprecation warnings
		panic(err)                                                  // force end program
	}
	err = db.Ping()

	if err != nil {
		logs.Log().Warn("Cant connect to DB!")
	}
	return &MySQLClient{db}
}

// ViewStats return sql.DBStats data
func (c *MySQLClient) ViewStats() sql.DBStats {
	return c.Stats
}
