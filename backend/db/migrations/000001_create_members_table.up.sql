CREATE TABLE IF NOT EXISTS members(
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    middle_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    dob TIMESTAMP NOT NULL,
    gender  VARCHAR(100) NOT NULL,
    id_no VARCHAR(20),
    id_issue_date TIMESTAMP,
    id_issue_place VARCHAR(100),
    phone VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    --description TEXT,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER REFERENCES users(id),
    updated_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by INTEGER REFERENCES users(id)
);
