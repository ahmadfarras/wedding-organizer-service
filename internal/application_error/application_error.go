package application_error

import "errors"

var (
	NotFoundErr = errors.New("Resource Not Found")
)
