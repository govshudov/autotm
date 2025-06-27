package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type PostgreSQLCartRepository struct {
	db *sql.DB
}

func NewPostgreSQLCartRepository(db *sql.DB) (*PostgreSQLCartRepository, error) {
	return &PostgreSQLCartRepository{db: db}, nil
}
func (p *PostgreSQLCartRepository) Test(ctx context.Context) error {
	fmt.Println("it is repository")
	return nil
}
