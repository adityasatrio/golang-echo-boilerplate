package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

// Table shape matches ent/schema/role_user.go (entimport-generated; DO NOT EDIT there).
// Ent automigrate is disabled (configs/database/connection_sqlent.go); this Goose
// migration is the executed source of truth for DDL, ent/schema/*.go documents the
// entity shape for code generation only. No drift: verified against ent/migrate/schema.go.
func init() {
	goose.AddMigrationContext(upAddTableForRoleUsers, downAddTableForRoleUsers)
}

func upAddTableForRoleUsers(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx,
		`CREATE TABLE role_users (
            id bigint unsigned NOT NULL AUTO_INCREMENT,
            versions bigint NOT NULL,
            created_by varchar(255) NOT NULL,
            created_at datetime NOT NULL,
            updated_by varchar(255) NULL,
            updated_at datetime NOT NULL,
            deleted_by varchar(255) NULL,
            deleted_at datetime NULL,
            user_id bigint unsigned NULL,
            role_id bigint unsigned NULL,
            PRIMARY KEY (id)
        );
    `)
	return err
}

func downAddTableForRoleUsers(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE role_users")
	return err
}
