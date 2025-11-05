package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
}
type ErrorResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
}
type ValidationErrorResponse struct {
	Status int `json:"status"`
	Errors []string `json:"errors"`
}

func SendResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Response{
		Status: statusCode,
		Data: data,
	})
}

func SendError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		Status: statusCode,
		Message: message,
	})
}

func SendValidationErrorResponse(w http.ResponseWriter, errs validator.ValidationErrors) {
	errors := make([]string, 0)
	for _, err := range errs {
		switch err.ActualTag(){
		case "required":
			errors = append(errors, fmt.Sprintf("Field %s is required", err.Field()))
		default:
			errors = append(errors, fmt.Sprintf("Field %s is not valid", err.Field()))
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ValidationErrorResponse{
		Status: http.StatusBadRequest,
		Errors: errors,
	})
}