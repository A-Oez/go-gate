CREATE TABLE log_inbound (
    id SERIAL PRIMARY KEY,
    request_id TEXT NOT NULL,
    time TIMESTAMP NOT NULL,
    method TEXT NOT NULL,
    uri TEXT NOT NULL,
    remote TEXT NOT NULL,
    headers JSONB,
    body TEXT
);