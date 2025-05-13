package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"
	"time"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    var err error

    for i := 0; i < 10; i++ {
        DB, err = sql.Open("postgres", dsn)
        if err == nil && DB.Ping() == nil {
            log.Println("✅ Conexión a PostgreSQL exitosa")
            return
        }
        log.Println("Esperando base de datos...")
        time.Sleep(2 * time.Second)
    }

    log.Fatal("No se pudo conectar a la base de datos después de varios intentos")
}
