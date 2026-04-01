package repository

import "myapp/internal/model"

type HealthRepository interface {
	GetHealth() model.HealthResponse
}

type healthRepository struct{}

func NewHealthRepository() HealthRepository {
	return &healthRepository{}
}

func (r *healthRepository) GetHealth() model.HealthResponse {
	return model.HealthResponse{Status: "ok"}
}
