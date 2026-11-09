package global

import "errors"

var (
	ERR_INVALID_TYPE             = errors.New("ERROR: invalid Type of T")
	ERR_DATABASE_TOO_MUCH        = errors.New("ERROR: func returnet to much data")
	ERR_KEY_NOT_FOUND            = errors.New("ERROR: key wasnt found!")
	ERR_FIELD_NOT_FOUND          = errors.New("ERROR: field not found!")
	ERR_ENTRY_EXPIRED            = errors.New("ERROR: entry expired!")
	ERR_SECRET_EXTRACTION_FAILED = errors.New("ERROR: failed to extract secret!")
	ERR_PATH_EMPTY               = errors.New("ERROR: path empty")
	ERR_TENANT_ID_INVALID        = errors.New("ERROR: TenantId invalid")
	ERR_STRING_EMTPY             = errors.New("ERROR: strin is empty")
)
