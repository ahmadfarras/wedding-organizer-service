package http

type ResponseError string

const (
	ForbiddenError        ResponseError = "FORBIDDEN"
	ResourceNotFoundError ResponseError = "RESOURCE_NOT_FOUND"
	InternalServerError   ResponseError = "INTERNAL_SERVER_ERROR"
)
