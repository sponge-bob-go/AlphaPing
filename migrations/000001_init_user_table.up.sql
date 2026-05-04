CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    tg_id VARCHAR(50) UNIQUE,
    tg_username VARCHAR(100)
);