CREATE TABLE sessions (
    id SERIAL PRIMARY KEY NOT NULL,
    user_email varchar(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP
);