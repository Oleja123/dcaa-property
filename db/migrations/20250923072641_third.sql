-- +goose Up
SELECT 'up SQL query';
ALTER TABLE properties
ADD COLUMN property_name VARCHAR(255);

-- +goose Down
SELECT 'down SQL query';
ALTER TABLE properties
DROP COLUMN property_name;