package service

import "forward/internal/repository"

type Service struct {
	Auth *AuthService
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Auth: newAuthService(repo.Auth),
	}
}
