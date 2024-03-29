package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddTableForExample, downAddTableForExample)
}

func upAddTableForExample(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.ExecContext(ctx,
		`CREATE TABLE example (
            id bigint unsigned NOT NULL AUTO_INCREMENT,
            username VARCHAR(50) NOT NULL,
            email VARCHAR(100) NOT NULL,
            PRIMARY KEY (id)
        );
    `)
	return err
}

func downAddTableForExample(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.ExecContext(ctx, "DROP TABLE example")
	return err
}
