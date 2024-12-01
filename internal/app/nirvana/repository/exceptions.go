package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"nirvana/internal/app/nirvana/usecases/model"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		db: pool,
	}
}

func (r *Repository) CreateException(ctx context.Context, name string, attributes model.ExceptionAttributes) error {
	query := `INSERT INTO exceptions (name, attributes) VALUES ($1, $2)`

	attributesJSON, err := json.Marshal(attributes)
	if err != nil {
		return fmt.Errorf("error marshaling attributes: %w", err)
	}

	_, err = r.db.Exec(ctx, query, name, attributesJSON)
	if err != nil {
		return fmt.Errorf("failed to insert exception: %w", err)
	}
	return nil
}

func (r *Repository) CheckException(ctx context.Context, name string, attributes model.ExceptionAttributes) (bool, error) {
	query := `SELECT COUNT(*) FROM exceptions WHERE name = $1 AND attributes = $2`

	attributesJSON, err := json.Marshal(attributes)
	if err != nil {
		return false, fmt.Errorf("error marshaling attributes: %w", err)
	}

	var count int
	err = r.db.QueryRow(ctx, query, name, attributesJSON).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check exception: %w", err)
	}

	return count > 0, nil
}
