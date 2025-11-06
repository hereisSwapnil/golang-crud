package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/hereisSwapnil/golang-crud/internal/config"
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
	result, err := s.Db.Exec("INSERT INTO students (name, age, email) VALUES (?, ?, ?)", name, age, email)
	if err != nil {
		return 0, fmt.Errorf("failed to create student: %v", err)
	}
	return result.LastInsertId()
}