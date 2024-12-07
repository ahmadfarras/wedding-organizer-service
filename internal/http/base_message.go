package http

type ResponseMessage string

const (
	SuccessMessage ResponseMessage = "Success"

	ForbiddenErrorMessage           ResponseMessage = "Doesn't have sufficient access permission"
	ResourceNotFoundErrorMessage    ResponseMessage = "Resource Not Found"
	InternalServerErrorErrorMessage ResponseMessage = "Internal Server Error"
)
