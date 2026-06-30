package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

// Table shape matches ent/schema/user.go (entimport-generated; DO NOT EDIT there).
// Ent automigrate is disabled (configs/database/connection_sqlent.go); this Goose
// migration is the executed source of truth for DDL, ent/schema/*.go documents the
// entity shape for code generation only. No drift: verified against ent/migrate/schema.go.
func init() {
	goose.AddMigrationContext(upAddTableForUsers, downAddTableForUsers)
}

func upAddTableForUsers(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx,
		`CREATE TABLE users (
            id bigint unsigned NOT NULL AUTO_INCREMENT,
            versions bigint NOT NULL,
            created_by varchar(255) NOT NULL,
            created_at datetime NOT NULL,
            updated_by varchar(255) NULL,
            updated_at datetime NOT NULL,
            deleted_by varchar(255) NULL,
            deleted_at datetime NULL,
            name varchar(255) NOT NULL,
            password varchar(255) NOT NULL,
            avatar varchar(255) NOT NULL,
            role_id bigint unsigned NOT NULL,
            is_verified tinyint(1) NOT NULL,
            email varchar(255) NULL,
            email_verified_at datetime NULL,
            remember_token varchar(255) NULL,
            social_media_id varchar(255) NULL,
            login_type varchar(255) NULL,
            sub_specialist varchar(255) NULL,
            firebase_token varchar(255) NULL,
            info varchar(255) NULL,
            description varchar(255) NULL,
            specialist varchar(255) NULL,
            phone varchar(255) NULL,
            last_access_at datetime NULL,
            pregnancy_mode tinyint(1) NULL,
            latest_skip_update datetime NULL,
            latest_deleted_at datetime NULL,
            PRIMARY KEY (id),
            UNIQUE KEY users_email_key (email),
            UNIQUE KEY users_phone_key (phone)
        );
    `)
	return err
}

func downAddTableForUsers(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE users")
	return err
}
