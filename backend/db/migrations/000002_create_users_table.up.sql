CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    password TEXT,
    role VARCHAR(100),
    enabled bool,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
