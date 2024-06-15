-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS vinyls (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    artist VARCHAR(255) NOT NULL,
    releasedate SMALLINT NOT NULL,
    price SMALLINT NOT NULL,
    rating FLOAT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS vinyls;
-- +goose StatementEnd
