CREATE TABLE sessions (
    id UUID PRIMARY KEY, 
    user_email varchar(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP
);