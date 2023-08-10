package db

import (
	"database/sql"
	config2 "github.com/farhanmvr/go-editor/config"
	"github.com/farhanmvr/go-editor/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"log"
	"time"
)

var db *sql.DB

const (
	maxRetries           = 3
	defaultRetryInterval = 3
)

func InitDB() {
	config := config2.GetConfig()
	var err error
	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("mysql", config.DB.Conn)
		if err != nil {
			log.Printf("Error connecting to database: %v. Retrying...", err)
			time.Sleep(time.Second * defaultRetryInterval)
		} else {
			break
		}
	}
	if err != nil {
		log.Fatal("Failed to connect to database after retries")
	}

	for i := 0; i < maxRetries; i++ {
		err = db.Ping()
		if err != nil {
			log.Printf("Error ping to database: %v. Retrying...", err)
			time.Sleep(time.Second * defaultRetryInterval)
		} else {
			break
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}

func PerformMigrations() {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	migrationsDir := "./db/migrations"
	sourceDriver, err := (&file.File{}).Open(migrationsDir)
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("file", sourceDriver, "mysql", driver)
	if err != nil {
		log.Fatal(err)
	}

	// DB migration up
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return db
}

func InsertCodeSnippet(code, name, status string) (*models.CodeSnippet, error) {
	id := uuid.New().String()
	query := "INSERT INTO code_snippets (id,name,code,status) VALUES (?,?,?,?)"
	_, err := db.Exec(query, id, name, code, status)
	if err != nil {
		return nil, err
	}
	return GetCodeSnippetById(id)
}

func GetAllCodeSnippets() ([]*models.CodeSnippet, error) {
	rows, err := db.Query("SELECT * FROM code_snippets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*models.CodeSnippet{}
	for rows.Next() {
		var snippet models.CodeSnippet
		var dateCreated []uint8

		if err := rows.Scan(&snippet.ID, &snippet.Name, &snippet.Code, &snippet.Status, &dateCreated); err != nil {
			return nil, err
		}

		// Convert the []uint8 value to a *time.Time value
		snippet.DateCreated, err = time.Parse("2006-01-02 15:04:05", string(dateCreated))
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, &snippet)
	}

	return snippets, nil
}

func GetCodeSnippetById(id string) (*models.CodeSnippet, error) {
	codeSnippet := models.CodeSnippet{}

	// Query the inserted row using its ID
	rquery := "SELECT * FROM code_snippets WHERE id = ?"
	var dateCreated []uint8
	err := db.QueryRow(rquery, id).Scan(
		&codeSnippet.ID, &codeSnippet.Name, &codeSnippet.Code, &codeSnippet.Status, &dateCreated,
	)
	// Convert the []uint8 value to a *time.Time value
	codeSnippet.DateCreated, err = time.Parse("2006-01-02 15:04:05", string(dateCreated))
	if err != nil {
		return &codeSnippet, err
	}

	return &codeSnippet, nil
}

func DeleteSnippetById(id string) (int64, error) {
	query := "DELETE FROM code_snippets WHERE id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
