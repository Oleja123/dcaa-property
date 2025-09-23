-- +goose Up
SELECT 'up SQL query';
CREATE TABLE properties (
    id BIGSERIAL PRIMARY KEY,
    addr TEXT NOT NULL,
    price INT,
    info TEXT,
    category_id INT NOT NULL,
    last_update TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);

-- +goose Down
SELECT 'down SQL query';
DROP TABLE properties;
