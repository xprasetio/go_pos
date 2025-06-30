package schemas

const SchemaOutlets = `
CREATE TABLE IF NOT EXISTS outlets (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		merchant_id BIGINT,
		name VARCHAR(255),
		province VARCHAR(255),
		city VARCHAR(255),
		address TEXT,
		phone VARCHAR(255),
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL,
		open_time VARCHAR(255),
		close_time VARCHAR(255),
		close_order BOOLEAN,
		job_id_open VARCHAR(255),
		job_id_close VARCHAR(255),
		area_province_id BIGINT DEFAULT 21,
		INDEX idx_outlets_area_province_id (area_province_id),
		INDEX idx_outlets_merchant_id (merchant_id)
	)
`
