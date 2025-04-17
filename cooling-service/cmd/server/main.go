package main

import (
    "cooling-service/internal/api"
    "net/http"
)

func main() {
    mux := api.SetupRouter()
    http.ListenAndServe(":8080", mux)
}