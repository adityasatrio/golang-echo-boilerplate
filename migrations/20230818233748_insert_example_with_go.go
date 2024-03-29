package migrations

import (
	"context"
	"database/sql"
	"errors"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upInsertExampleWithGo, downInsertExampleWithGo)
}

func upInsertExampleWithGo(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	id, err := getExampleID(tx, "James")
	if err != nil {
		return err
	}
	if id == 0 {
		query := "INSERT INTO example (username, email) VALUES (?, ?)"
		if _, err := tx.ExecContext(ctx, query, "James", "james@email.com"); err != nil {
			return err
		}
	}
	return nil
}

func downInsertExampleWithGo(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	query := "DELETE FROM example WHERE username = ?"
	if _, err := tx.ExecContext(ctx, query, "James"); err != nil {
		return err
	}
	return nil
}

func getExampleID(db *sql.Tx, username string) (int, error) {
	var id int
	err := db.QueryRow("SELECT id FROM example WHERE username = ?", username).Scan(&id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}
	return id, nil
}
