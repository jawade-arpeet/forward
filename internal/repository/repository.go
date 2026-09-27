package repository

import "forward/internal/client"

type Repository struct{}

func New(clt *client.Client) *Repository {
	return &Repository{}
}
