package cooling

import "time"

type CoolingPeriod struct {
    CreditID      string
    ClientID      string
    ContractTime  time.Time
    CoolingHours  int
    PrincipalPaid bool
    Withdrawn     bool
    Status        string
    UpdatedAt     time.Time
}