package api

import (
    "cooling-service/internal/cooling"
    "cooling-service/internal/storage"
    "net/http"
)

var store *storage.PostgresStorage

func InitRouter(s *storage.PostgresStorage) *http.ServeMux {
    store = s
    mux := http.NewServeMux()
    mux.HandleFunc("/cooling/register", cooling.MakeRegisterHandler(store))
    mux.HandleFunc("/cooling/validate", cooling.MakeValidateHandler(store))
    mux.HandleFunc("/cooling/pay", cooling.MakePayHandler(store))
    mux.HandleFunc("/cooling/withdraw", cooling.MakeWithdrawHandler(store))
    mux.HandleFunc("/cooling/status", cooling.MakeStatusHandler(store))
    mux.HandleFunc("/cooling/report", cooling.MakeReportHandler(store))
    return mux
}