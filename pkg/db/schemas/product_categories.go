package schemas

const SchemaProductCategories = `
CREATE TABLE IF NOT EXISTS product_categories (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255),
		show_in_pos BOOLEAN DEFAULT TRUE,
		outlet_id BIGINT,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL,
		INDEX idx_product_categories_outlet_id (outlet_id)
	)
`
