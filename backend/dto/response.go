package dto

import "github.com/Rhtymn/eduall-case-study/backend/apperror"

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func ResponseError(appErr *apperror.AppError) Response {
	return Response{
		Message: appErr.Message(),
		Data:    nil,
	}
}

func ResponseOK(data any) Response {
	return Response{
		Message: "ok",
		Data:    data,
	}
}
