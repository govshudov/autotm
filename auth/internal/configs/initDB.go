package configs

import (
	"database/sql"
	"fmt"
)

const (
	defaultIdleConns    = 5
	defaultMaxOpenConns = 10
)

func InitDb(configs *Config) (*sql.DB, error) {
	var err error
	dbHost := configs.Storage.Psql.Host
	dbPort := configs.Storage.Psql.Port
	dbUser := configs.Storage.Psql.Username
	dbPassword := configs.Storage.Psql.Password
	dbName := configs.Storage.Psql.Database
	connectTimeOut := 5

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=%d",
		dbHost, dbPort, dbUser, dbPassword, dbName, connectTimeOut)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(defaultIdleConns)
	db.SetMaxOpenConns(defaultMaxOpenConns)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
