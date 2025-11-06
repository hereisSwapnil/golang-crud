package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/hereisSwapnil/golang-crud/internal/config"
	"github.com/hereisSwapnil/golang-crud/internal/types"

	// as we are just using sqlite3, we need to import it not directly using it
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(config *config.Config)(*Sqlite, error){
	db, err := sql.Open("sqlite3", config.StoragePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL,
		email TEXT NOT NULL
	)`)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create table: %v", err)
	}
	return &Sqlite{Db: db}, nil
}

func (s *Sqlite) CreateStudent(name string, age int, email string) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO students (name, age, email) VALUES (?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(name, age, email)
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %v", err)
	}
	return result.LastInsertId()
}

func (s *Sqlite) GetStudent(id int) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM students WHERE id = ?")
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Age, &student.Email)
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to query row: %v", err)
	}
	return student, nil
}

func (s *Sqlite) GetAllStudents() ([]types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM students")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("failed to query statement: %v", err)
	}
	defer rows.Close()
	var students []types.Student
	for rows.Next() {
		var student types.Student
		err = rows.Scan(&student.Id, &student.Name, &student.Age, &student.Email)
		if err != nil {
			return nil, fmt.Errorf("failed to scan student: %v", err)
		}
		students = append(students, student)
	}
	return students, nil
}