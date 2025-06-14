package handlers
import (
    "database/sql"
    "log"
    "net/http"
)

func CreateListHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req struct {
            Title      string `json:"title"`
            Category   string `json:"category"`
            AutoValidate bool `json:"autoValidate"`
        }
        json.NewDecoder(r.Body).Decode(&req)

        _, err := db.Exec(`
            INSERT INTO lists (user_id, title, category, auto_validate) 
            VALUES (?, ?, ?, ?)`, 1, req.Title, req.Category, req.AutoValidate)
        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        w.WriteHeader(http.StatusCreated)
    }
}
