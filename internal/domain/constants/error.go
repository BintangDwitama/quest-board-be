package constants

var GeneralErrorConcrete *GeneralError

type GeneralError struct {
	Code    int
	Message string
}

func NewGeneralError(code int, message string) *GeneralError {
	return &GeneralError{
		Code:    code,
		Message: message,
	}
}

func (e *GeneralError) Error() string {
	return e.Message
}

func (e *GeneralError) ErrorCode() int {
	return e.Code
}
