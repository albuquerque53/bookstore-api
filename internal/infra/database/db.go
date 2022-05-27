package database

import (
	"context"
	"database/sql"
)

type DatabaseManager struct {
	DB *sql.DB
}

func NewDatabaseManager() *DatabaseManager {
	return &DatabaseManager{
		DB: ConectToDatabase(),
	}
}

func (dm *DatabaseManager) Query(ctx context.Context, q string, args ...interface{}) (*sql.Rows, error) {
	defer dm.DB.Close()

	stmt, err := dm.DB.Prepare(q)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return stmt.QueryContext(ctx, args...)
}

func (dm *DatabaseManager) QueryRow(ctx context.Context, q string, args ...interface{}) (*sql.Row, error) {
	defer dm.DB.Close()

	stmt, err := dm.DB.Prepare(q)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return stmt.QueryRowContext(ctx, args...), nil
}
