package student

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/hereisSwapnil/golang-crud/internal/storage"
	"github.com/hereisSwapnil/golang-crud/internal/types"
	response "github.com/hereisSwapnil/golang-crud/internal/utils"
)

func New(storage storage.Storage) http.HandlerFunc {
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

		id, err := storage.CreateStudent(student.Name, student.Age, student.Email)
		if err != nil {
			response.SendError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create student: %v", err))
			return
		}

		fmt.Println("Student created successfully", id)
		response.SendResponse(w, http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}
}

func Get(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.SendError(w, http.StatusBadRequest, "Invalid student ID")
			return
		}
		student, err := storage.GetStudent(idInt)
		if err != nil {
			response.SendError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get student: %v", err))
			return
		}
		response.SendResponse(w, http.StatusOK, student)
	}
}

func GetAll(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		students, err := storage.GetAllStudents()
		if err != nil {
			response.SendError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get students: %v", err))
			return
		}
		response.SendResponse(w, http.StatusOK, students)
	}
}