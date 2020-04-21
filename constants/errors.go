package constants

// 0xx - General errors
const (
	ErrNotImplemented = 000
)

// 3xx - Authorization errors
const (
	ErrMissingToken          = 300
	ErrInvalidToken          = 301
	ErrInvalidInvitationCode = 302
	ErrUsernameExists        = 303
	ErrPasswordHashFailed    = 304
	ErrInvalidPassword       = 305
	ErrCreateTokenFailed     = 306
	ErrCompliancedNotPassed  = 307
)

// 1xx - DB error
const (
	ErrDbConn           = 100
	ErrDbNoRowsReturned = 101
	ErrDbLookupFailed   = 102
	ErrDbCreateFailed   = 103
	ErrDbUpdateFailed   = 104
	ErrDbDeleteFailed   = 105

	ErrDbCacheLookupFailed = 110
	ErrDbCacheSetFailed    = 111
)

// 4xx - User request (validation) error
const (
	ErrRequestBadJSON          = 400
	ErrRequestValidationFailed = 401
	ErrMissingParam            = 402
	ErrDuplicateResource       = 403
	ErrInvalidImage            = 404
	ErrRequestBadFormData      = 405
	ErrInvalidFile             = 406
	ErrInvalidFileType         = 407
	ErrResourceNotFound        = 408
)

// 5xx - Internal Error
const (
	ErrInternalServerError  = 500
	ErrUnauthorizedResource = 501
)

// 6xx - 3rd party error
const (
	ErrS3UploadFailed = 601
)
