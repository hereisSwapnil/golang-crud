package student

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/hereisSwapnil/golang-crud/internal/types"
	response "github.com/hereisSwapnil/golang-crud/internal/utils"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			response.SendError(w, http.StatusBadRequest, "Failed to decode student data")
			return
		}

		if err := validator.New().Struct(student); err != nil {
			response.SendValidationErrorResponse(w, err.(validator.ValidationErrors))
			return
		}

		fmt.Println("Student created successfully", student)
		response.SendResponse(w, http.StatusOK, student)
	}
}