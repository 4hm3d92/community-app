CREATE TABLE IF NOT EXISTS payments(
    id SERIAL PRIMARY KEY,
    payment_type INTEGER,
	receipt_no INTEGER UNIQUE,
	member_id INTEGER REFERENCES members(id),
	amount INTEGER,
	added_on TIMESTAMP,
	added_by INTEGER REFERENCES users(id),
	paid_on TIMESTAMP,
	updated_by INTEGER REFERENCES users(id),
	updated_on TIMESTAMP
);