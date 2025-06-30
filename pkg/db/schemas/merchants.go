package schemas

const SchemaMerchants = `
CREATE TABLE IF NOT EXISTS merchants (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		code VARCHAR(255),
		name VARCHAR(255),
		merchants_app VARCHAR(255),
		merchants_key VARCHAR(255),
		base_url VARCHAR(255),
		prefix_trx_id VARCHAR(255),
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL)
`
