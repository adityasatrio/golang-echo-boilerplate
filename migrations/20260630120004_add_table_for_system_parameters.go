package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

// Table shape matches ent/schema/system_parameter.go (entimport-generated; DO NOT EDIT there).
// Ent automigrate is disabled (configs/database/connection_sqlent.go); this Goose
// migration is the executed source of truth for DDL, ent/schema/*.go documents the
// entity shape for code generation only. No drift: verified against ent/migrate/schema.go.
func init() {
	goose.AddMigrationContext(upAddTableForSystemParameters, downAddTableForSystemParameters)
}

func upAddTableForSystemParameters(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx,
		`CREATE TABLE system_parameters (
            id int NOT NULL AUTO_INCREMENT,
            versions bigint NOT NULL,
            created_by varchar(255) NOT NULL,
            created_at datetime NOT NULL,
            updated_by varchar(255) NULL,
            updated_at datetime NOT NULL,
            deleted_by varchar(255) NULL,
            deleted_at datetime NULL,
            ` + "`key`" + ` varchar(255) NOT NULL,
            value varchar(255) NOT NULL,
            PRIMARY KEY (id),
            UNIQUE KEY system_parameters_key_key (` + "`key`" + `),
            KEY systemparameter_key (` + "`key`" + `),
            KEY systemparameter_value (value)
        );
    `)
	return err
}

func downAddTableForSystemParameters(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE system_parameters")
	return err
}
