package simple_sql_example

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	INSERT INTO  public.tasks (name, description, completed, created_at)
	VALUES ('выгулять джека', 'погулять c джеком в 4', FALSE, '2026-10-30 15:26');
	`
	_, err := conn.Exec(ctx, sqlQuery)
	
	return err
}
