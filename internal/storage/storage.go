package storage

import "github.com/hereisSwapnil/golang-crud/internal/types"

type Storage interface {
	CreateStudent(name string, age int, email string) (int64, error)
	GetStudent(id int) (types.Student, error)
	GetAllStudents() ([]types.Student, error)
	UpdateStudent(id int, name string, age int, email string) error
	DeleteStudent(id int) error
}