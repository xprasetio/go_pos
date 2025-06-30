package schemas

const SchemaUsers = `
CREATE TABLE IF NOT EXISTS users (
	id BIGINT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(255),
	password VARCHAR(255),
	phone VARCHAR(255),
	enabled BOOLEAN,
	merchant_id BIGINT,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
	is_employee BOOLEAN DEFAULT TRUE,
	email VARCHAR(255),
	current_portion INT DEFAULT 0,
	avatar VARCHAR(255),
	access_token VARCHAR(255),
	INDEX idx_users_merchant_id (merchant_id)
)`
