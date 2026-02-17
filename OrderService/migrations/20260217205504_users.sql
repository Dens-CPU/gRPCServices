-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS  users(
id SERIAL PRIMARY KEY,
user_id INT NOT NULL UNIQUE,
created_at TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
