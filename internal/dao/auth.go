package dao

import "uuid"

type AccountInfo struct {
	ID              uuid.UUID
	Email           string
	PasswordHash    string
	IsActive        bool
	IsEmailVerified bool
}
