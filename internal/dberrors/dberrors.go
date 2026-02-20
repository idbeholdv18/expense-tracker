package dberrors

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	SqlStateUniqueViolation = "23505"
	SqlStateForeignKeyViolation = "23503"
)

func IsUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError

	return errors.As(err, &pgErr) && pgErr.Code == SqlStateUniqueViolation
}

func IsForeignKeyViolation(err error) bool {
	var pgErr *pgconn.PgError

	return errors.As(err, &pgErr) && pgErr.Code == SqlStateForeignKeyViolation
}
