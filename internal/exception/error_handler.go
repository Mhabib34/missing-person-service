package exception

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Mhbib34/missing-person-service/internal/dto"
	"github.com/Mhbib34/missing-person-service/internal/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ErrorHandler(ctx *gin.Context, err any) {

	if notFoundError(ctx, err) {
		return
	}

	if validationErrors(ctx, err) {
		return
	}

	if conflictError(ctx, err) {
		return
	}

	internalServerError(ctx, err)
}

func validationErrors(ctx *gin.Context, err any) bool {
	ex, ok := err.(validator.ValidationErrors)
	if ok {

		webResponse := dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Error:  ex.Error(),
		}

		helper.WriteToResponseBody(ctx, http.StatusBadRequest, webResponse)
		return true
	}
	return false
}


func notFoundError(ctx *gin.Context, err any) bool {
	if e, ok := err.(error); ok {
			
		if errors.Is(e, gorm.ErrRecordNotFound) {
			webResponse := dto.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Error:  "Report not found",
			}

			helper.WriteToResponseBody(ctx, http.StatusNotFound, webResponse)
			return true
		}
	}
	return false
}

func conflictError(ctx *gin.Context, err any) bool {
	ex, ok := err.(ConflictError)
	if ok {

		webResponse := dto.WebResponse{
			Code:   http.StatusConflict,
			Status: "CONFLICT",
			Error:  ex.Error(),
		}

		helper.WriteToResponseBody(ctx, http.StatusConflict, webResponse)
		return true
	}
	return false
}


func internalServerError(ctx *gin.Context, err any) {

	webResponse := dto.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Error:  fmt.Sprintf("%v", err),
	}

	helper.WriteToResponseBody(ctx, http.StatusInternalServerError, webResponse)
}