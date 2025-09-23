-- +goose Up
SELECT 'up SQL query';
ALTER TABLE properties
ALTER COLUMN property_name SET NOT NULL,
ALTER COLUMN last_update SET NOT NULL;

-- +goose Down
SELECT 'down SQL query';
ALTER TABLE properties
ALTER COLUMN property_name DROP NOT NULL,
ALTER COLUMN last_update DROP NOT NULL;
