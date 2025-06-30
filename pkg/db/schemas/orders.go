package schemas

const SchemaOrders = `
CREATE TABLE IF NOT EXISTS orders (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		trx_no VARCHAR(255),
		trx_date DATETIME,
		user_id BIGINT,
		disc_percentage FLOAT,
		disc_value FLOAT,
		tax_percentage FLOAT DEFAULT 1.1,
		tax_value FLOAT,
		outlet_id BIGINT,
		total FLOAT DEFAULT 0.0,
		subtotal FLOAT DEFAULT 0.0,
		status VARCHAR(255) DEFAULT "waiting",
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL,
		ordered_by VARCHAR(255),
		discount_id BIGINT,
		disc_name VARCHAR(255),
		processed_by INT,
		INDEX idx_orders_discount_id (discount_id),
		INDEX idx_orders_outlet_id (outlet_id),
		INDEX idx_orders_user_id (user_id)
	)
`
