CREATE TABLE IF NOT EXISTS WarehousesTable (
    id SERIAL PRIMARY KEY,
    identifier VARCHAR(255) NOT NULL,
    addr VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Products (
    id SERIAL PRIMARY KEY,
    identifier VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    weight VARCHAR(50),
    barcode VARCHAR(255) UNIQUE
);

CREATE TABLE IF NOT EXISTS product_key_values (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    key INTEGER NOT NULL,
    value TEXT NOT NULL,
    CONSTRAINT pk_product_key_value UNIQUE (product_id, key)
);

CREATE TABLE IF NOT EXISTS Inventory (
    id SERIAL PRIMARY KEY,
    warehouse_id VARCHAR(255) NOT NULL,
    product_id VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 0,
    price NUMERIC(10, 2) NOT NULL DEFAULT 0.00,
    discount NUMERIC(5, 2) NOT NULL DEFAULT 0.00,
    CONSTRAINT uc_warehouse_product UNIQUE (warehouse_id, product_id)
);

CREATE TABLE IF NOT EXISTS Analytics (
    id SERIAL PRIMARY KEY,
    warehouse_id VARCHAR(255) NOT NULL,
    product_id VARCHAR(255) NOT NULL,
    sold_goods INTEGER NOT NULL DEFAULT 0,
    total_sum NUMERIC(15, 2) NOT NULL DEFAULT 0.00,
    report_date DATE NOT NULL DEFAULT CURRENT_DATE,
    CONSTRAINT uc_analytics_record UNIQUE (warehouse_id, product_id, report_date)
);