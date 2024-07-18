CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        order_number VARCHAR(100) NOT NULL UNIQUE,
                        user_id INT NOT NULL,
                        total_amount DECIMAL(10,2) NOT NULL,
                        order_status VARCHAR(50) NOT NULL,
                        order_date TIMESTAMP NOT NULL,
                        delivery_date TIMESTAMP,
                        shipping_address TEXT NOT NULL,
                        created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                        updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
                        deleted_at TIMESTAMP
);

CREATE TABLE order_items (
                             id SERIAL PRIMARY KEY,
                             order_id INT NOT NULL,
                             product_id INT NOT NULL,
                             quantity INT NOT NULL,
                             unit_price DECIMAL(10,2) NOT NULL,
                             created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                             updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
                             deleted_at TIMESTAMP,
                             FOREIGN KEY (order_id) REFERENCES orders(id),
                             FOREIGN KEY (product_id) REFERENCES products(id)
);
