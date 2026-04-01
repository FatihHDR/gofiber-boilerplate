package service

import (
	"myapp/internal/model"
	"myapp/internal/repository"
)

type HealthService interface {
	Check() model.HealthResponse
}

type healthService struct {
	repo repository.HealthRepository
}

func NewHealthService(repo repository.HealthRepository) HealthService {
	return &healthService{repo: repo}
}

func (s *healthService) Check() model.HealthResponse {
	return s.repo.GetHealth()
}
