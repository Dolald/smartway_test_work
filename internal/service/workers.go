package service

import "github.com/Dolald/smartway_test_work/internal/repository"

type WorkersService struct {
	repository repository.Repository
}

func NewWorkersService(repository *repository.Repository) *WorkersService {
	return &WorkersService{repository: *repository}
}
