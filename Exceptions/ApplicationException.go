package Exceptions

import "fmt"

type ApplicationException struct {
	Code    int
	Message string
	Detail  error
}

func (e *ApplicationException) Error() string {
	return fmt.Sprintf("Code %v: %v", e.Code, e.Message)
}
func NewApplicationException(code int, message string, detail *error) *ApplicationException {
	return &ApplicationException{
		Code:    code,
		Message: message,
		Detail:  *detail,
	}
}
