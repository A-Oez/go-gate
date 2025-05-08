CREATE TABLE admin_users (
    id SERIAL PRIMARY KEY,                
    email VARCHAR(255) NOT NULL UNIQUE,    
    password_hash TEXT NOT NULL,           
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO admin_users (email, password_hash)
VALUES 
('admin@example.com', '$2a$12$rG4snDm.xEbetwVu/o2I5.mw4C0dvNlFJQn8QZe/mdhXoMQAEUztq'); --admin123

