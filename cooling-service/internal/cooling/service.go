package cooling

import (
    "encoding/json"
    "net/http"
    "sync"
    "time"
)

var coolingData = make(map[string]*CoolingPeriod)
var lock = sync.Mutex{}

type RegisterRequest struct {
    CreditID     string    `json:"credit_id"`
    ClientID     string    `json:"client_id"`
    ContractTime time.Time `json:"contract_time"`
    CoolingHours int       `json:"cooling_hours"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var req RegisterRequest
    json.NewDecoder(r.Body).Decode(&req)

    lock.Lock()
    defer lock.Unlock()

    coolingData[req.CreditID] = &CoolingPeriod{
        CreditID:     req.CreditID,
        ClientID:     req.ClientID,
        ContractTime: req.ContractTime,
        CoolingHours: req.CoolingHours,
        PrincipalPaid: false,
        Withdrawn:     false,
        Status:        "not_requested",
        UpdatedAt:     time.Now(),
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Registered"))
}

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
    creditID := r.URL.Query().Get("credit_id")

    lock.Lock()
    defer lock.Unlock()

    cp, exists := coolingData[creditID]
    if !exists {
        http.Error(w, "Credit not found", http.StatusNotFound)
        return
    }

    expiry := cp.ContractTime.Add(time.Duration(cp.CoolingHours) * time.Hour)
    isValid := time.Now().Before(expiry)

    response := map[string]interface{}{
        "is_valid":   isValid,
        "expires_at": expiry.Format(time.RFC3339),
        "message":    "Cooling period is " + map[bool]string{true: "active", false: "expired"}[isValid],
    }

    json.NewEncoder(w).Encode(response)
}

func PayHandler(w http.ResponseWriter, r *http.Request) {
    creditID := r.URL.Query().Get("credit_id")

    lock.Lock()
    defer lock.Unlock()

    cp, exists := coolingData[creditID]
    if !exists {
        http.Error(w, "Credit not found", http.StatusNotFound)
        return
    }

    cp.PrincipalPaid = true
    cp.UpdatedAt = time.Now()
    w.Write([]byte("Principal marked as paid"))
}

func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
    creditID := r.URL.Query().Get("credit_id")

    lock.Lock()
    defer lock.Unlock()

    cp, exists := coolingData[creditID]
    if !exists {
        http.Error(w, "Credit not found", http.StatusNotFound)
        return
    }

    expiry := cp.ContractTime.Add(time.Duration(cp.CoolingHours) * time.Hour)

    if time.Now().After(expiry) {
        json.NewEncoder(w).Encode(map[string]interface{}{
            "status":  "rejected",
            "message": "Cooling period expired",
        })
        return
    }

    if !cp.PrincipalPaid {
        json.NewEncoder(w).Encode(map[string]interface{}{
            "status":  "rejected",
            "message": "Principal not paid",
        })
        return
    }

    cp.Withdrawn = true
    cp.Status = "annulled"
    cp.UpdatedAt = time.Now()

    json.NewEncoder(w).Encode(map[string]interface{}{
        "status":  "accepted",
        "message": "Credit agreement annulled",
    })
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
    creditID := r.URL.Query().Get("credit_id")

    lock.Lock()
    defer lock.Unlock()

    cp, exists := coolingData[creditID]
    if !exists {
        http.Error(w, "Credit not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(map[string]interface{}{
        "credit_id":  cp.CreditID,
        "status":     cp.Status,
        "updated_at": cp.UpdatedAt.Format(time.RFC3339),
    })
}

func ReportHandler(w http.ResponseWriter, r *http.Request) {
    lock.Lock()
    defer lock.Unlock()

    var report []map[string]interface{}
    for _, cp := range coolingData {
        entry := map[string]interface{}{
            "credit_id":     cp.CreditID,
            "client_id":     cp.ClientID,
            "status":        cp.Status,
            "contract_time": cp.ContractTime.Format(time.RFC3339),
            "updated_at":    cp.UpdatedAt.Format(time.RFC3339),
        }
        report = append(report, entry)
    }

    json.NewEncoder(w).Encode(report)
}