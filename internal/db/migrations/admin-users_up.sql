CREATE TABLE admin_users (
    id SERIAL PRIMARY KEY,                
    email VARCHAR(255) NOT NULL UNIQUE,    
    password_hash TEXT NOT NULL,           
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO admin_users (email, password_hash)
VALUES 
('admin@example.com', '$2a$10$93eqbLnM0JlzXOzv.kG/a/KnFEwY5.Lq4sToOp.EpDe/iV1PCgm/C'); --admin123

