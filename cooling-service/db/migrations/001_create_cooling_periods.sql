CREATE TABLE IF NOT EXISTS cooling_periods (
    credit_id TEXT PRIMARY KEY,
    client_id TEXT NOT NULL,
    contract_time TIMESTAMP NOT NULL,
    cooling_hours INTEGER NOT NULL,
    principal_paid BOOLEAN NOT NULL DEFAULT FALSE,
    withdrawn BOOLEAN NOT NULL DEFAULT FALSE,
    status TEXT NOT NULL DEFAULT 'not_requested',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);