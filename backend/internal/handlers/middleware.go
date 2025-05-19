package handlers

import(
	// "log"
	"net/http"
	"os"

	// "github.com/joho/godotenv"
)

func EnableCORS(next http.Handler) http.Handler {
  	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		frontendOrigin := os.Getenv("FRONTEND_URL")
		w.Header().Set("Access-Control-Allow-Origin", frontendOrigin)
    	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    	if r.Method == http.MethodOptions {
      	w.WriteHeader(http.StatusNoContent)
      	return
    	}

    	next.ServeHTTP(w, r)
  	})
}
