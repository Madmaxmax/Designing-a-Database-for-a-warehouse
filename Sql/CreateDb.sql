CREATE TABLE IF NOT EXISTS orders(
    order_id INTEGER PRIMARY KEY,
    customer_name VARCHAR(255),
    date TIMESTAMP DEFAULT NOW() NOT NULL
);
CREATE TABLE IF NOT EXISTS order_details (
    order_id INTEGER NOT NULL PRIMARY KEY REFERENCES orders(order_id),
    order_info JSONB NOT NULL
);
CREATE TABLE IF NOT EXISTS shelving(
    shelf_id INTEGER PRIMARY KEY,
    shelf_name VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS products(
    product_id INTEGER,
    shelf_id INTEGER NOT NULL REFERENCES shelving(shelf_id),
    is_main BOOLEAN NOT NULL,
    product_count INTEGER NOT NULL,
    product_name VARCHAR(255) NOT NULL
);