-- +goose Up
CREATE TABLE grids(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    grid TEXT NOT NULL
);

-- +goose Down
DROP table grids;