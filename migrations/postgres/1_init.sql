-- +goose Up
CREATE TABLE exceptions (
                            id SERIAL PRIMARY KEY,
                            name VARCHAR(255) NOT NULL,
                            attributes JSONB NOT NULL,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE exceptions;