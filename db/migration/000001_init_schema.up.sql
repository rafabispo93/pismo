BEGIN;

CREATE TABLE accounts(
    id SERIAL PRIMARY KEY,
    balance FLOAT NOT NULL DEFAULT 0,
    document_number TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE transactions(
    id SERIAL PRIMARY KEY,
    account_id SERIAL NOT NULL REFERENCES accounts(id),
    transaction_type TEXT NOT NULL,
    amount FLOAT NOT NULL,
    event_date timestamptz NOT NULL DEFAULT (now())
);

COMMIT;