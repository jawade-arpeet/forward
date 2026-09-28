package errs

import "errors"

var (
	ErrPgNoRows              = errors.New("no rows")
	ErrPgUniqueViolation     = errors.New("unique violation")
	ErrPgForeignKeyViolation = errors.New("foreign key violation")
	ErrPgNotNullViolation    = errors.New("not null violation")
	ErrPgCheckViolation      = errors.New("check violation")
)
