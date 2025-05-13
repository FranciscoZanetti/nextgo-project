package main

import (
    "log"
    "net/http"

    "github.com/FranciscoZanetti/nextgo-project/backend/internal/handlers"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(`{"status":"ok"}`))
    })

    http.HandleFunc("/tasks", handlers.HandleTasks)
    http.HandleFunc("/tasks/", handlers.HandleTaskByID)

    log.Println("Servidor escuchando en http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
