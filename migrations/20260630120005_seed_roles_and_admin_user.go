package migrations

import (
	"context"
	"database/sql"
	"errors"
	"github.com/pressly/goose/v3"
	"golang.org/x/crypto/bcrypt"
)

// bootstrapAdminPassword is a placeholder credential for the seeded admin
// account. Rotate it immediately in any non-local environment.
const bootstrapAdminPassword = "ChangeMe123!"

// Seeds rows for roles/users tables defined in ent/schema/role.go and
// ent/schema/user.go (entimport-generated; DO NOT EDIT there). No drift:
// verified against ent/migrate/schema.go.
func init() {
	goose.AddMigrationContext(upSeedRolesAndAdminUser, downSeedRolesAndAdminUser)
}

func upSeedRolesAndAdminUser(ctx context.Context, tx *sql.Tx) error {
	adminRoleID, err := getRoleID(tx, "Admin")
	if err != nil {
		return err
	}
	if adminRoleID == 0 {
		if _, err := tx.ExecContext(ctx,
			"INSERT INTO roles (id, versions, created_by, created_at, updated_at, name, text) VALUES (1, 1, 'system', NOW(), NOW(), ?, ?)",
			"Admin", "Full access to all pages and resources"); err != nil {
			return err
		}
	}

	userRoleID, err := getRoleID(tx, "User")
	if err != nil {
		return err
	}
	if userRoleID == 0 {
		if _, err := tx.ExecContext(ctx,
			"INSERT INTO roles (id, versions, created_by, created_at, updated_at, name, text) VALUES (2, 1, 'system', NOW(), NOW(), ?, ?)",
			"User", "Read-only access to system parameters"); err != nil {
			return err
		}
	}

	adminUserID, err := getUserIDByEmail(tx, "admin@example.com")
	if err != nil {
		return err
	}
	if adminUserID == 0 {
		hashed, err := bcrypt.GenerateFromPassword([]byte(bootstrapAdminPassword), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		result, err := tx.ExecContext(ctx,
			`INSERT INTO users (versions, created_by, created_at, updated_at, name, password, avatar, role_id, is_verified, email)
             VALUES (1, 'system', NOW(), NOW(), ?, ?, '', 1, 1, ?)`,
			"Admin", string(hashed), "admin@example.com")
		if err != nil {
			return err
		}

		newUserID, err := result.LastInsertId()
		if err != nil {
			return err
		}

		if _, err := tx.ExecContext(ctx,
			"INSERT INTO role_users (versions, created_by, created_at, updated_at, user_id, role_id) VALUES (1, 'system', NOW(), NOW(), ?, 1)",
			newUserID); err != nil {
			return err
		}
	}

	return nil
}

func downSeedRolesAndAdminUser(ctx context.Context, tx *sql.Tx) error {
	if _, err := tx.ExecContext(ctx, "DELETE FROM role_users WHERE user_id IN (SELECT id FROM users WHERE email = ?)", "admin@example.com"); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, "DELETE FROM users WHERE email = ?", "admin@example.com"); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, "DELETE FROM roles WHERE name IN (?, ?)", "Admin", "User"); err != nil {
		return err
	}
	return nil
}

func getRoleID(db *sql.Tx, name string) (int64, error) {
	var id int64
	err := db.QueryRow("SELECT id FROM roles WHERE name = ?", name).Scan(&id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}
	return id, nil
}

func getUserIDByEmail(db *sql.Tx, email string) (int64, error) {
	var id int64
	err := db.QueryRow("SELECT id FROM users WHERE email = ?", email).Scan(&id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}
	return id, nil
}
