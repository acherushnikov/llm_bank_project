package api

import (
    "cooling-service/internal/cooling"
    "cooling-service/internal/storage"
    "log"
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var store *storage.PostgresStorage
var validate = validator.New()

func InitRouter(s *storage.PostgresStorage) *mux.Router {
    store = s
    r := mux.NewRouter()
	r.Use(loggingMiddleware)

	r.HandleFunc("/cooling/register", withValidation(cooling.MakeRegisterHandler(store))).Methods("POST")
	r.HandleFunc("/cooling/validate", withValidation(cooling.MakeValidateHandler(store))).Methods("POST")
	r.HandleFunc("/cooling/pay", withValidation(cooling.MakePayHandler(store))).Methods("POST")
	r.HandleFunc("/cooling/withdraw", withValidation(cooling.MakeWithdrawHandler(store))).Methods("POST")
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

func withValidation(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type bodyReader interface {
			Validate() error
		}

		var req bodyReader
		if err := parseAndValidate(r, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func parseAndValidate(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(v); err != nil {
		return err
	}
	return validate.Struct(v)
}
