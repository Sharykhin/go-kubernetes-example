-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE test ADD COLUMN first_name VARCHAR(80);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE test DROP COLUMN first_name;