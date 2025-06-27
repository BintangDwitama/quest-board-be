package constants

type standardResponse struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func SuccessResponse(message string, data interface{}) standardResponse {
	return standardResponse{
		Message: message,
		Data: data,
	}
}

func FailedResponse(message string, detail interface{}) standardResponse {
	return standardResponse{
		Message: message,
		Data: detail,
	}
}