package schemas

const SchemaRoles = `
CREATE TABLE IF NOT EXISTS roles (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		code VARCHAR(255),
		name VARCHAR(255),
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL)
`
