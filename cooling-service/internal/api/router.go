package api

import (
    "cooling-service/internal/cooling"
    "net/http"
)

func SetupRouter() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/cooling/register", cooling.RegisterHandler)
    mux.HandleFunc("/cooling/validate", cooling.ValidateHandler)
    mux.HandleFunc("/cooling/pay", cooling.PayHandler)
    mux.HandleFunc("/cooling/withdraw", cooling.WithdrawHandler)
    mux.HandleFunc("/cooling/status", cooling.StatusHandler)
    mux.HandleFunc("/cooling/report", cooling.ReportHandler)
    return mux
}