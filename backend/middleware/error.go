package middleware

import (
	"net/http"

	"github.com/Rhtymn/eduall-case-study/backend/apperror"
	"github.com/Rhtymn/eduall-case-study/backend/dto"
	"github.com/gin-gonic/gin"
)

func NewErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		err := ctx.Errors[0].Err

		apperr, ok := err.(*apperror.AppError)
		if ok {
			ctx.AbortWithStatusJSON(
				getHttpStatus(apperr),
				dto.ResponseError(apperr),
			)
			return
		}

		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			dto.ResponseError(apperror.Wrap(err).(*apperror.AppError)),
		)
	}
}

func getHttpStatus(err *apperror.AppError) int {
	switch err.Code() {
	case apperror.CodeInvalidValidation:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
