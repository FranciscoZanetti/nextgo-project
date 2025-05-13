package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "github.com/FranciscoZanetti/nextgo-project/backend/internal/database"
    "github.com/FranciscoZanetti/nextgo-project/backend/internal/models"
)

func HandleTasks(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        rows, err := database.DB.Query("SELECT id, title, description, created_at FROM task")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var tasks []models.Task
        for rows.Next() {
            var t models.Task
            if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            tasks = append(tasks, t)
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(tasks)

    case http.MethodPost:
        var t models.Task
        if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        err := database.DB.QueryRow(
            "INSERT INTO task (title, description) VALUES ($1, $2) RETURNING id, created_at",
            t.Title, t.Description,
        ).Scan(&t.ID, &t.CreatedAt)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(t)

    default:
        http.Error(w, "Método no soportado", http.StatusMethodNotAllowed)
    }
}

func HandleTaskByID(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case http.MethodGet:
        var t models.Task
        err := database.DB.QueryRow("SELECT id, title, description, created_at FROM task WHERE id = $1", id).
            Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt)

        if err != nil {
            http.Error(w, "No encontrado", http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(t)

    case http.MethodPut:
        var t models.Task
        if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        _, err := database.DB.Exec("UPDATE task SET title=$1, description=$2 WHERE id=$3", t.Title, t.Description, id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusNoContent)

    case http.MethodDelete:
        _, err := database.DB.Exec("DELETE FROM task WHERE id=$1", id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusNoContent)

    default:
        http.Error(w, "Método no soportado", http.StatusMethodNotAllowed)
    }
}
