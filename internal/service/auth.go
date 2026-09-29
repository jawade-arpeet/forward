package service

import (
	"context"
	"errors"
	"forward/internal/argon2id"
	"forward/internal/errs"
	"forward/internal/repository"
	"strings"
)

type AuthService struct {
	authRepo *repository.AuthRepository
}

func newAuthService(authRepo *repository.AuthRepository) *AuthService {
	return &AuthService{authRepo: authRepo}
}

func (s *AuthService) SignUp(
	ctx context.Context,
	email string,
	password string,
) error {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)
	password = strings.TrimSpace(password)

	passwordHash, err := argon2id.Hash(password)
	if err != nil {
		return errs.Internal(err)
	}

	_, err = s.authRepo.CreateAccount(ctx, email, passwordHash)
	if err != nil {
		if errors.Is(err, errs.ErrPgUniqueViolation) {
			return errs.Conflict(
				errs.CodeAccountAlreadyExists,
				"account already exists",
				err,
			)
		}

		return errs.Internal(err)
	}

	return nil
}

func (s *AuthService) SignIn(
	ctx context.Context,
	email string,
	password string,
) (string, error) {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)
	password = strings.TrimSpace(password)

	acc, err := s.authRepo.GetAccountByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, errs.ErrPgNoRows) {
			return "", errs.Unauthorized(
				errs.CodeInvalidEmailOrPassword,
				"invalid email or password",
				nil,
			)
		}

		return "", errs.Internal(err)
	}

	ok, err := argon2id.Compare(password, acc.PasswordHash)
	if err != nil {
		return "", errs.Internal(err)
	}

	if !ok {
		return "", errs.Unauthorized(
			errs.CodeInvalidEmailOrPassword,
			"invalid email or password",
			nil,
		)
	}

	return acc.ID.String(), nil
}
