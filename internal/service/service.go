package service

import "forward/internal/repository"

type Service struct{}

func New(repo *repository.Repository) *Service {
	return &Service{}
}
