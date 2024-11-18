package service

import (
	"github.com/Dolald/smartway_test_work/internal/repository"
)

type Workers interface {
}

type Service struct {
	Workers
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Workers: NewWorkersService(repos),
	}
}
