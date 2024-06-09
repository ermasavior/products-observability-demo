package utils

import "net/http"

const MessageInternalServerError = "Internal server error"

type Response struct {
	Status HeaderResponse `json:"status"`
	Data   interface{}    `json:"data,omitempty"`
	Meta   interface{}    `json:"meta,omitempty"`
}

type HeaderResponse struct {
	Code               int      `json:"code"`
	Message            string   `json:"message"`
	ErrorDetailMessage []string `json:"errors,omitempty"`
}

func NewFailedResponse(errorCode int, message string) Response {
	return Response{
		Status: HeaderResponse{
			Code:    errorCode,
			Message: message,
		},
	}
}

func NewSuccessResponse(data interface{}) Response {
	return Response{
		Status: HeaderResponse{
			Code:    http.StatusOK,
			Message: "success",
		},
		Data: data,
	}
}
