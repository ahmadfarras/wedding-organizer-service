package http

type ResponseCode string

const (
	SuccessCode ResponseCode = "00"

	CredentialErrorCode     ResponseCode = "40"
	ForbiddenErrorCode      ResponseCode = "41"
	NotFoundErrorCode       ResponseCode = "42"
	InternalServerErrorCode ResponseCode = "43"
)
