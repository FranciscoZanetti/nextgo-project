package main

import (
    "log"
    "net/http"
	"os"
    "time"

    "github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
    "github.com/FranciscoZanetti/nextgo-project/backend/internal/database"
	"github.com/FranciscoZanetti/nextgo-project/backend/internal/handlers"
)

func main() {
    err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("Sentry initialization failed: %v\n", err)
	}
	defer sentry.Flush(2 * time.Second)

	database.InitDB()

	port := os.Getenv("PORT")
	url := os.Getenv("URL")

    sentryHandler := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})

    mux := http.NewServeMux()

    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(`{"status":"ok"}`))
    })

    mux.Handle("/tasks", handlers.EnableCORS(http.HandlerFunc(handlers.HandleTasks)))
    mux.Handle("/tasks/", handlers.EnableCORS(http.HandlerFunc(handlers.HandleTaskByID)))

    log.Printf("Server running on %s:%s", url, port)
    log.Fatal(http.ListenAndServe(":" + port, mux))
}
