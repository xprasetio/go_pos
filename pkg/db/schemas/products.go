package schemas

const SchemaProducts = `
CREATE TABLE IF NOT EXISTS products (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		product_category_id BIGINT,
		name VARCHAR(255),
		sku_number VARCHAR(255),
		description TEXT,
		het FLOAT,
		hpp FLOAT,
		hp FLOAT,
		image VARCHAR(255),
		parent_id INT,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL,
		INDEX idx_products_name (name),
		INDEX idx_products_parent_id (parent_id),
		INDEX idx_products_product_category_id (product_category_id)
	)
`
