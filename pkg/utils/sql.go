package utils

import (
	"errors"
	"github.com/jackc/pgconn"
)

func IsDuplicateEntryError(err error) bool {
	var pgError *pgconn.PgError
	return errors.As(err, &pgError) && pgError.Code == "23505"
}
