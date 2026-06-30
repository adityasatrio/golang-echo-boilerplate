-- Seed default roles (Admin, User).
-- Idempotent: safe to run multiple times against either MySQL or PostgreSQL
-- (the active db.configs.driver). Run after migrations have created the
-- `roles` table.

INSERT INTO roles (versions, created_by, created_at, updated_at, name, text)
SELECT 1, 'system', NOW(), NOW(), 'Admin', 'Full access to all pages and resources'
WHERE NOT EXISTS (SELECT 1 FROM roles WHERE name = 'Admin');

INSERT INTO roles (versions, created_by, created_at, updated_at, name, text)
SELECT 1, 'system', NOW(), NOW(), 'User', 'Read-only access to system parameters'
WHERE NOT EXISTS (SELECT 1 FROM roles WHERE name = 'User');
