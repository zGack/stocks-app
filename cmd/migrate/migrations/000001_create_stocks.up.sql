CREATE TABLE IF NOT EXISTS stock (
    id SERIAL PRIMARY KEY,
    ticker VARCHAR(255) NOT NULL,
    target_from VARCHAR(255) NOT NULL,
    target_to VARCHAR(255) NOT NULL,
    company VARCHAR(255) NOT NULL,
    action VARCHAR(255) NOT NULL,
    brokerage VARCHAR(255) NOT NULL,
    rating_from VARCHAR(255) NOT NULL,
    rating_to VARCHAR(255) NOT NULL,
    time TIMESTAMP NULL
);
