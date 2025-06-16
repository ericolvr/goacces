-- Criação da tabela users para SQLite
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    mobile TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role INTEGER NOT NULL,
    status BOOLEAN NOT NULL
);

-- Inserção de dados iniciais
INSERT INTO users (name, mobile, password, role, status) VALUES
    ('Admin', '11999999999', 'admin123', 1, 1),
    ('João Silva', '11988888888', 'senha123', 2, 1),
    ('Maria Souza', '11977777777', 'senha456', 2, 0);
