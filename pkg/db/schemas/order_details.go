package schemas

const SchemaOrderDetails = `
CREATE TABLE IF NOT EXISTS order_details (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		order_id BIGINT,
		product_id BIGINT,
		add_on_id BIGINT,
		qty INT DEFAULT 0,
		price FLOAT DEFAULT 0.0,
		total FLOAT DEFAULT 0.0,
		note TEXT,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL,
		INDEX idx_order_details_add_on_id (add_on_id),
		INDEX idx_order_details_order_id (order_id),
		INDEX idx_order_details_product_id (product_id)
	)
`
