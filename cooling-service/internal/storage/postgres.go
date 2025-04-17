package storage

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/lib/pq"
)

type PostgresStorage struct {
    DB *sql.DB
}

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

func NewPostgresStorage(host, port, user, password, dbname string) (*PostgresStorage, error) {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    log.Println("✅ Подключение к PostgreSQL успешно установлено.")
    return &PostgresStorage{DB: db}, nil
}

func (s *PostgresStorage) CreateCoolingPeriod(cp CoolingPeriod) error {
    _, err := s.DB.Exec(`
        INSERT INTO cooling_periods (
            credit_id, client_id, contract_time, cooling_hours,
            principal_paid, withdrawn, status, updated_at
        ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
    `, cp.CreditID, cp.ClientID, cp.ContractTime, cp.CoolingHours,
        cp.PrincipalPaid, cp.Withdrawn, cp.Status, cp.UpdatedAt)
    return err
}

func (s *PostgresStorage) GetCoolingPeriod(creditID string) (*CoolingPeriod, error) {
    row := s.DB.QueryRow(`SELECT credit_id, client_id, contract_time, cooling_hours, principal_paid,
        withdrawn, status, updated_at FROM cooling_periods WHERE credit_id = $1`, creditID)

    var cp CoolingPeriod
    err := row.Scan(&cp.CreditID, &cp.ClientID, &cp.ContractTime, &cp.CoolingHours,
        &cp.PrincipalPaid, &cp.Withdrawn, &cp.Status, &cp.UpdatedAt)

    if err != nil {
        return nil, err
    }
    return &cp, nil
}

func (s *PostgresStorage) UpdateCoolingPeriod(cp CoolingPeriod) error {
    _, err := s.DB.Exec(`
        UPDATE cooling_periods SET
            principal_paid = $1,
            withdrawn = $2,
            status = $3,
            updated_at = $4
        WHERE credit_id = $5
    `, cp.PrincipalPaid, cp.Withdrawn, cp.Status, cp.UpdatedAt, cp.CreditID)
    return err
}

func (s *PostgresStorage) ListCoolingPeriods() ([]CoolingPeriod, error) {
    rows, err := s.DB.Query(`SELECT credit_id, client_id, contract_time, cooling_hours, principal_paid,
        withdrawn, status, updated_at FROM cooling_periods`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var cps []CoolingPeriod
    for rows.Next() {
        var cp CoolingPeriod
        err := rows.Scan(&cp.CreditID, &cp.ClientID, &cp.ContractTime, &cp.CoolingHours,
            &cp.PrincipalPaid, &cp.Withdrawn, &cp.Status, &cp.UpdatedAt)
        if err != nil {
            return nil, err
        }
        cps = append(cps, cp)
    }
    return cps, nil
}
func (s *PostgresStorage) RunMigrations() error {
    query := `
    CREATE TABLE IF NOT EXISTS cooling_periods (
        credit_id TEXT PRIMARY KEY,
        client_id TEXT NOT NULL,
        contract_time TIMESTAMP NOT NULL,
        cooling_hours INTEGER NOT NULL,
        principal_paid BOOLEAN NOT NULL DEFAULT FALSE,
        withdrawn BOOLEAN NOT NULL DEFAULT FALSE,
        status TEXT NOT NULL DEFAULT 'not_requested',
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );`
    _, err := s.DB.Exec(query)
    return err
}