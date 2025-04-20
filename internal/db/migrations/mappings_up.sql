-- up.sql

CREATE TABLE mappings (
    id SERIAL PRIMARY KEY,
    method VARCHAR(10),
    public_path VARCHAR(255),
    service_scheme VARCHAR(10),
    service_host VARCHAR(255),
    service_path VARCHAR(255)
);

-- Optional: Füge einige Anfangsdaten hinzu
INSERT INTO mappings (method, public_path, service_scheme, service_host, service_path)
VALUES
('GET', '/api', 'http', 'apitest:8080', '/test');