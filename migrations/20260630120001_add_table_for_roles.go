package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddTableForRoles, downAddTableForRoles)
}

func upAddTableForRoles(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx,
		`CREATE TABLE roles (
            id bigint unsigned NOT NULL AUTO_INCREMENT,
            versions bigint NOT NULL,
            created_by varchar(255) NOT NULL,
            created_at datetime NOT NULL,
            updated_by varchar(255) NULL,
            updated_at datetime NOT NULL,
            deleted_by varchar(255) NULL,
            deleted_at datetime NULL,
            name varchar(255) NOT NULL,
            text varchar(255) NOT NULL,
            PRIMARY KEY (id)
        );
    `)
	return err
}

func downAddTableForRoles(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE roles")
	return err
}
