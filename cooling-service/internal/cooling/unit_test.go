package cooling

import (
    "testing"
    "time"
)

func TestCoolingPeriodExpiry(t *testing.T) {
    cp := CoolingPeriod{
        CreditID:     "CR123",
        ClientID:     "CL123",
        ContractTime: time.Now().Add(-2 * time.Hour),
        CoolingHours: 1,
    }

    expiry := cp.ContractTime.Add(time.Duration(cp.CoolingHours) * time.Hour)
    if time.Now().Before(expiry) {
        t.Error("Expected cooling period to be expired")
    }
}

func TestCoolingPeriodActive(t *testing.T) {
    cp := CoolingPeriod{
        CreditID:     "CR456",
        ClientID:     "CL456",
        ContractTime: time.Now(),
        CoolingHours: 2,
    }

    expiry := cp.ContractTime.Add(time.Duration(cp.CoolingHours) * time.Hour)
    if !time.Now().Before(expiry) {
        t.Error("Expected cooling period to be active")
    }
}