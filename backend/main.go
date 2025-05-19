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

    mux := http.NewServeMux()

    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(`{"status":"ok"}`))
    })

    mux.Handle("/tasks", handlers.EnableCORS(http.HandlerFunc(handlers.HandleTasks)))
    mux.Handle("/tasks/", handlers.EnableCORS(http.HandlerFunc(handlers.HandleTaskByID)))

    log.Printf("Server running on %s:%s", url, port)
    log.Fatal(http.ListenAndServe(":" + port, mux))
}
