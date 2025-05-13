package main

import (
    "log"
    "net/http"
	"os"

    "github.com/FranciscoZanetti/nextgo-project/backend/internal/database"
	"github.com/FranciscoZanetti/nextgo-project/backend/internal/handlers"
)

func main() {
	database.InitDB()

	port := os.Getenv("PORT")
	url := os.Getenv("URL")

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(`{"status":"ok"}`))
    })

    http.HandleFunc("/tasks", handlers.HandleTasks)
    http.HandleFunc("/tasks/", handlers.HandleTaskByID)

    log.Printf("Servidor escuchando en %s:%s", url, port)
    log.Fatal(http.ListenAndServe(":" + port, nil))
}
