CREATE TABLE IF NOT EXISTS Warehouses (
    id SERIAL PRIMARY KEY,
    warehouseId VARCHAR(255) NOT NULL,
    addr VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Products (
    id SERIAL PRIMARY KEY,
    productId VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    weight NUMERIC(10, 2) NOT NULL DEFAULT 0.00,
    barcode VARCHAR(255) UNIQUE
);

CREATE TABLE IF NOT EXISTS ProductKeyValues (
    productId VARCHAR(255) NOT NULL,
    key VARCHAR(255) NOT NULL,
    value VARCHAR(255) NOT NULL,
    PRIMARY KEY (productId, key)
);

CREATE TABLE IF NOT EXISTS Inventory (
    id SERIAL PRIMARY KEY,
    warehouseId VARCHAR(255) NOT NULL,
    productId VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 0,
    price NUMERIC(10, 2) NOT NULL DEFAULT 0.00,
    discount NUMERIC(5, 2) NOT NULL DEFAULT 0.00,
    CONSTRAINT ucWarehouseProduct UNIQUE (warehouseId, productId)
);

CREATE TABLE IF NOT EXISTS Analytics (
    id SERIAL PRIMARY KEY,
    warehouseId VARCHAR(255) NOT NULL,
    productId VARCHAR(255) NOT NULL,
    soldGoods INTEGER NOT NULL DEFAULT 0,
    totalSum NUMERIC(15, 2) NOT NULL DEFAULT 0.00,
    CONSTRAINT ucAnalyticsRecord UNIQUE (
        warehouseId,
        productId
    )
);