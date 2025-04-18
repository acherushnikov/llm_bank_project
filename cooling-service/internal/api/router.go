package api

import (
    "cooling-service/internal/cooling"
    "cooling-service/internal/storage"
    "net/http"
    "encoding/json"
	"log"
)

var store *storage.PostgresStorage

func InitRouter(s *storage.PostgresStorage) *mux.Router {
    store = s
    r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/cooling/register", cooling.MakeRegisterHandler(store)).Methods("POST")
	r.HandleFunc("/cooling/validate", cooling.MakeValidateHandler(store)).Methods("POST")
	r.HandleFunc("/cooling/pay", cooling.MakePayHandler(store)).Methods("POST")
	r.HandleFunc("/cooling/withdraw", cooling.MakeWithdrawHandler(store)).Methods("POST")
	r.HandleFunc("/cooling/status", cooling.MakeStatusHandler(store)).Methods("GET")
	r.HandleFunc("/cooling/report", cooling.MakeReportHandler(store)).Methods("GET")
    return r
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}