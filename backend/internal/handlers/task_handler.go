package handlers

import (
    "encoding/json"
    "net/http"
    "strings"

    "github.com/FranciscoZanetti/nextgo-project/backend/internal/models"
)

func HandleTasks(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        tasks := []models.Task{
            {ID: "1", Title: "Learn Go", Completed: false, CreatedAt: "2024-01-01T00:00:00Z"},
            {ID: "2", Title: "Build API", Completed: true, CreatedAt: "2024-01-02T00:00:00Z"},
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(tasks)

    case http.MethodPost:
        var newTask models.Task
        if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        newTask.ID = "123" // ID simulado
        newTask.CreatedAt = "2024-01-10T00:00:00Z"
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(newTask)

    default:
        http.Error(w, "Método no soportado", http.StatusMethodNotAllowed)
    }
}

func HandleTaskByID(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/tasks/")

    switch r.Method {
    case http.MethodGet:
        task := models.Task{ID: id, Title: "Mock Task", Completed: false, CreatedAt: "2024-01-05T00:00:00Z"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(task)

    case http.MethodPut:
        var updatedTask models.Task
        if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        updatedTask.ID = id
        updatedTask.CreatedAt = "2024-01-05T00:00:00Z"
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(updatedTask)

    case http.MethodDelete:
        w.WriteHeader(http.StatusNoContent)

    default:
        http.Error(w, "Método no soportado", http.StatusMethodNotAllowed)
    }
}
