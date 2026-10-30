package simple_sql_example

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error{
	sqlQuert := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		name VARCHAR(200) NOT NULL,
		description VARCHAR(1000) NOT NULL,
		completed BOOLEAN NOT NULL,
		created_at TIMESTAMP NOT NULL,
		completed_at TIMESTAMP,

		UNIQUE(name)
	);
	`

	 _, err := conn.Exec(ctx,sqlQuert)

	 return err
}