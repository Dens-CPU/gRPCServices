-- Active: 1771190303876@@127.0.0.1@5432
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders(
    id SERIAL PRIMARY KEY,
    user_id  INT REFERENCES users(id),
    market_id INT REFERENCES markets(id),
    order_type TEXT NOT NULL,
    price FLOAT NOT NULL,
    quantity INT NOT NULL,
    status TEXT NOT NULL,
    order_id INT REFERENCES orders(id),
    Created_at TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd