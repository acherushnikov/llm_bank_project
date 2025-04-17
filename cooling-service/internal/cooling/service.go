package cooling

import (
    "cooling-service/internal/storage"
    "encoding/json"
    "net/http"
    "time"
)

// === Register ===
type RegisterRequest struct {
    CreditID     string    `json:"credit_id"`
    ClientID     string    `json:"client_id"`
    ContractTime time.Time `json:"contract_time"`
    CoolingHours int       `json:"cooling_hours"`
}

func MakeRegisterHandler(store *storage.PostgresStorage) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req RegisterRequest
        json.NewDecoder(r.Body).Decode(&req)

        cp := storage.CoolingPeriod{
            CreditID:      req.CreditID,
            ClientID:      req.ClientID,
            ContractTime:  req.ContractTime,
            CoolingHours:  req.CoolingHours,
            PrincipalPaid: false,
            Withdrawn:     false,
            Status:        "not_requested",
            UpdatedAt:     time.Now(),
        }

        err := store.CreateCoolingPeriod(cp)
        if err != nil {
            http.Error(w, "Ошибка создания записи", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        w.Write([]byte("Registered"))
    }
}

// === Validate ===
func MakeValidateHandler(store *storage.PostgresStorage) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        creditID := r.URL.Query().Get("credit_id")
        cp, err := store.GetCoolingPeriod(creditID)
        if err != nil {
            http.Error(w, "Кредит не найден", http.StatusNotFound)
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
}

// === Pay ===
func MakePayHandler(store *storage.PostgresStorage) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        creditID := r.URL.Query().Get("credit_id")
        cp, err := store.GetCoolingPeriod(creditID)
        if err != nil {
            http.Error(w, "Кредит не найден", http.StatusNotFound)
            return
        }

        cp.PrincipalPaid = true
        cp.UpdatedAt = time.Now()
        err = store.UpdateCoolingPeriod(*cp)
        if err != nil {
            http.Error(w, "Ошибка обновления", http.StatusInternalServerError)
            return
        }

        w.Write([]byte("Principal marked as paid"))
    }
}

// === Withdraw ===
func MakeWithdrawHandler(store *storage.PostgresStorage) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        creditID := r.URL.Query().Get("credit_id")
        cp, err := store.GetCoolingPeriod(creditID)
        if err != nil {
            http.Error(w, "Кредит не найден", http.StatusNotFound)
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
        err = store.UpdateCoolingPeriod(*cp)
        if err != nil {
            http.Error(w, "Ошибка обновления", http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(map[string]interface{}{
            "status":  "accepted",
            "message": "Credit agreement annulled",
        })
    }
}

// === Status ===
func MakeStatusHandler(store *storage.PostgresStorage) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        creditID := r.URL.Query().Get("credit_id")
        cp, err := store.GetCoolingPeriod(creditID)
        if err != nil {
            http.Error(w, "Кредит не найден", http.StatusNotFound)
            return
        }

        json.NewEncoder(w).Encode(map[string]interface{}{
            "credit_id":  cp.CreditID,
            "status":     cp.Status,
            "updated_at": cp.UpdatedAt.Format(time.RFC3339),
        })
    }
}

// === Report ===
func MakeReportHandler(store *storage.PostgresStorage) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        cps, err := store.ListCoolingPeriods()
        if err != nil {
            http.Error(w, "Ошибка выборки", http.StatusInternalServerError)
            return
        }

        var report []map[string]interface{}
        for _, cp := range cps {
            report = append(report, map[string]interface{}{
                "credit_id":     cp.CreditID,
                "client_id":     cp.ClientID,
                "status":        cp.Status,
                "contract_time": cp.ContractTime.Format(time.RFC3339),
                "updated_at":    cp.UpdatedAt.Format(time.RFC3339),
            })
        }

        json.NewEncoder(w).Encode(report)
    }
}