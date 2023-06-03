CREATE TABLE auth_user (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    last_login TIMESTAMPTZ DEFAULT current_timestamp,
    firstname VARCHAR(255),
    lastname VARCHAR(255),
    phone_number VARCHAR(255),
    email VARCHAR(255),
    is_superuser BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT current_timestamp,
    updated_at TIMESTAMPTZ DEFAULT current_timestamp
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    order_number VARCHAR(50) UNIQUE NOT NULL,
    status VARCHAR(20) NOT NULL,
    shipping_address VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    total DECIMAL DEFAULT 0,
    total_discount DECIMAL DEFAULT 0,
    is_cancelled BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT current_timestamp,
    updated_at TIMESTAMPTZ DEFAULT current_timestamp,
    FOREIGN KEY (user_id) REFERENCES auth_user (id)
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    code VARCHAR(20) UNIQUE NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    price DECIMAL DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT current_timestamp,
    updated_at TIMESTAMPTZ DEFAULT current_timestamp
);


CREATE TABLE order_items (
    ID SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT DEFAULT 0,
    price DECIMAL DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT current_timestamp,
    updated_at TIMESTAMPTZ DEFAULT current_timestamp,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);