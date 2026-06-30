-- Seed a bootstrap admin user with the Admin role.
-- Idempotent: safe to run multiple times against either MySQL or PostgreSQL
-- (the active db.configs.driver). Requires 0001_seed_default_roles.sql to
-- have run first.
--
-- Login: admin@example.com / ChangeMe123!
-- Rotate this password immediately in any non-local environment.

INSERT INTO users (versions, created_by, created_at, updated_at, name, password, avatar, role_id, is_verified, email)
SELECT 1, 'system', NOW(), NOW(), 'Admin',
       '$2a$10$kJ693kpseUIVbBUsoOL/juX/Qn//aaKaM3aWtWXxeyjCL.JdzXVIS',
       '', (SELECT id FROM roles WHERE name = 'Admin'), 1, 'admin@example.com'
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = 'admin@example.com');

INSERT INTO role_users (versions, created_by, created_at, updated_at, user_id, role_id)
SELECT 1, 'system', NOW(), NOW(), u.id, r.id
FROM users u, roles r
WHERE u.email = 'admin@example.com' AND r.name = 'Admin'
  AND NOT EXISTS (
    SELECT 1 FROM role_users ru WHERE ru.user_id = u.id AND ru.role_id = r.id
  );
