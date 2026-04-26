package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// postgres://postgres:135642@localhost:5432/postgres?sslmode=disable
func ConnetcToBD(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, "postgres://bogdanivanov:135642@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		return nil, err
	}

	return conn, nil

}
