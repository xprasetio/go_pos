package schemas

const SchemaUserRoles = `
CREATE TABLE IF NOT EXISTS user_roles (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		user_id BIGINT,
		role_id BIGINT,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL,
		INDEX idx_user_roles_role_id (role_id),
		INDEX idx_user_roles_user_id (user_id))
`
