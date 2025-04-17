package main

import (
    "cooling-service/internal/api"
    "cooling-service/internal/storage"
    "log"
    "net/http"
    "os"
)

func main() {
    db, err := storage.NewPostgresStorage(
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASS"),
        os.Getenv("DB_NAME"),
    )
    if err != nil {
        log.Fatal("DB error:", err)
    }

    if err := db.RunMigrations(); err != nil {
	log.Fatal("Migration failed:", err)
}

	mux := api.InitRouter(db)
    log.Println("🌐 Cooling Service API running on :8080")
    http.ListenAndServe(":8080", mux)
}