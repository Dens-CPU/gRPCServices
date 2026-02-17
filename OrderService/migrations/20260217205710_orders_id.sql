-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS  orders_id(
id SERIAL PRIMARY KEY,
orders_id INT NOT NULL UNIQUE,
created_at TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS  orders_id;
-- +goose StatementEnd
