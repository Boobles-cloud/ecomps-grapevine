package global

import "errors"

var (
	ERR_INVALID_TYPE      = errors.New("ERROR: invalid Type of T")
	ERR_DATABASE_TOO_MUCH = errors.New("ERROR: func returnet to much data")
)
