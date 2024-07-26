DROP DATABASE IF EXISTS eduall_db;
CREATE DATABASE eduall_db;
\c eduall_db;

CREATE TABLE products(
    id BIGSERIAL PRIMARY KEY,
    brand VARCHAR(50) NOT NULL,
    model VARCHAR,
    screen_size VARCHAR NOT NULL,
    color VARCHAR,
    harddisk VARCHAR(20),
    cpu VARCHAR,
    ram VARCHAR,
    os VARCHAR,
    special_features VARCHAR,
    graphics VARCHAR,
    graphic_coprocessor VARCHAR,
    cpu_speed VARCHAR,
    rating NUMERIC(2, 1),
    price NUMERIC(10, 2)
);