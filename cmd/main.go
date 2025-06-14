package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "modernc.org/sqlite"
)

type List struct {
	Title         string `json:"title"`
	Category      string `json:"category"`
	AutoValidate  bool   `json:"autoValidate"`
}

func main() {
	db, err := sql.Open("sqlite", "./top5.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := createTables(db); err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Post("/lists", createListHandler(db))

	log.Println("Server starting at :8080")
	http.ListenAndServe(":8080", r)
}

func createListHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var list List
		if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}

		_, err := db.Exec(`
			INSERT INTO lists (user_id, title, category, auto_validate) 
			VALUES (?, ?, ?, ?)`,
			1, list.Title, list.Category, list.AutoValidate)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("list created"))
	}
}

func createTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS lists (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		title TEXT,
		category TEXT,
		auto_validate BOOLEAN DEFAULT FALSE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(query)
	return err
}
