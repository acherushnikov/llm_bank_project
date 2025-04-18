package main

import (
    "cooling-service/internal/api"
    "cooling-service/internal/storage"
    "log"
    "net/http"
    "os"
    "os/signal"
	"syscall"
    "time"
)

func main() {
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASS")
    db_name := os.Getenv("DB_NAME")

    if host == "" 
    || port == ""
    || user == ""
    || pass == ""
    || db_name == "" {
		log.Fatal("Missing env variables")
	}

    db, err := storage.NewPostgresStorage(
        host,
        port,
        user,
        pass,
        db_name,
    )
    if err != nil {
        log.Fatal("DB error:", err)
    }
    defer db.Close()

    if err := db.RunMigrations(); err != nil {
        log.Fatal("Migration failed:", err)
    }

	mux := api.InitRouter(db)
    srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

    go func() {
		log.Println("Server starting on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

    quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed: %v", err)
	}
	log.Println("Server exited")
}