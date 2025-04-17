package main

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"

    "cooling-service/internal/api"
)

func TestCoolingServiceEndToEnd(t *testing.T) {
    router := api.SetupRouter()

    // 1. Register credit
    registerPayload := map[string]interface{}{
        "credit_id":     "CR001",
        "client_id":     "CL001",
        "contract_time": time.Now().Format(time.RFC3339),
        "cooling_hours": 24,
    }
    body, _ := json.Marshal(registerPayload)
    req := httptest.NewRequest("POST", "/cooling/register", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    if resp.Code != http.StatusCreated {
        t.Fatalf("Register failed, expected 201 got %d", resp.Code)
    }

    // 2. Validate cooling period
    req = httptest.NewRequest("GET", "/cooling/validate?credit_id=CR001", nil)
    resp = httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    var validateResp map[string]interface{}
    _ = json.Unmarshal(resp.Body.Bytes(), &validateResp)

    if !validateResp["is_valid"].(bool) {
        t.Fatalf("Expected cooling period to be valid")
    }

    // 3. Pay principal
    req = httptest.NewRequest("POST", "/cooling/pay?credit_id=CR001", nil)
    resp = httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    if resp.Code != http.StatusOK {
        t.Fatalf("Expected 200 for pay, got %d", resp.Code)
    }

    // 4. Withdraw request
    req = httptest.NewRequest("POST", "/cooling/withdraw?credit_id=CR001", nil)
    resp = httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    var withdrawResp map[string]interface{}
    _ = json.Unmarshal(resp.Body.Bytes(), &withdrawResp)

    if withdrawResp["status"] != "accepted" {
        t.Fatalf("Expected accepted, got %v", withdrawResp["status"])
    }

    // 5. Check status
    req = httptest.NewRequest("GET", "/cooling/status?credit_id=CR001", nil)
    resp = httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    var statusResp map[string]interface{}
    _ = json.Unmarshal(resp.Body.Bytes(), &statusResp)

    if statusResp["status"] != "annulled" {
        t.Fatalf("Expected status to be 'annulled', got %v", statusResp["status"])
    }
}