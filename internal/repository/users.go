package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type WorkersRepository struct {
	db *sqlx.DB
}

func NewWorkersRepository(db *sqlx.DB) *WorkersRepository {
	return &WorkersRepository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context) {

}
