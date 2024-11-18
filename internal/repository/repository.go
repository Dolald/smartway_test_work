package repository

import (
	"github.com/jmoiron/sqlx"
)

type Worker interface {
}

type Repository struct {
	Worker
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Worker: NewWorkersRepository(db),
	}
}
