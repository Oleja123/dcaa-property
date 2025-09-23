-- +goose Up
SELECT 'up SQL query';
ALTER TABLE properties
ALTER COLUMN price TYPE FLOAT;

-- +goose Down
SELECT 'down SQL query';
ALTER TABLE properties
ALTER COLUMN price TYPE INTEGER;