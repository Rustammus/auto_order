
-- +goose Up
-- SQL in this section is executed when the migration is applied



CREATE TABLE products (
                          id INTEGER PRIMARY KEY,
                          article TEXT NULL,
                          name TEXT NOT NULL,
                          description TEXT,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product_analog (
                                id INTEGER PRIMARY KEY,
                                product_id INTEGER NOT NULL,
                                analog_id INTEGER NOT NULL,
                                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                FOREIGN KEY (product_id) REFERENCES products(id),
                                FOREIGN KEY (analog_id) REFERENCES products(id)
);

CREATE TABLE product_suppliers (
                                   id INTEGER PRIMARY KEY,
                                   product_id INTEGER NOT NULL,
                                   sup_id INTEGER NOT NULL,
                                   sup_product_id INTEGER NULL,
                                   sup_code TEXT NULL,
                                   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                   FOREIGN KEY (product_id) REFERENCES products(id),
                                   FOREIGN KEY (sup_id) REFERENCES suppliers(id),
                                   UNIQUE(product_id, sup_id)  -- Одна связь продукт-поставщик
);

CREATE TABLE suppliers (
                           id INTEGER PRIMARY KEY,
                           name TEXT NOT NULL,
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE to_order (
                          id INTEGER PRIMARY KEY,
                          product_id INTEGER NOT NULL,
                          sup_id INTEGER NOT NULL,
                          sup_code TEXT NULL,
                          count INTEGER NOT NULL,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          FOREIGN KEY (product_id) REFERENCES products(id),
                          UNIQUE(product_id, sup_id)
);

CREATE INDEX idx_product_suppliers_product_id ON product_suppliers(product_id);
CREATE INDEX idx_product_suppliers_sup_id ON product_suppliers(sup_id);

CREATE INDEX idx_products_article ON products(article);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back