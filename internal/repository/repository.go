package repository

import "forward/internal/client"

type Repository struct {
	Auth *AuthRepository
}

func New(clt *client.Client) *Repository {
	return &Repository{
		Auth: newAuthRepository(clt.Postgres, clt.Redis),
	}
}
