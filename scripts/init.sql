CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    profile TEXT NOT NULL,
    document TEXT,
    pin INTEGER,
    coercion INTEGER,
    card_number INTEGER,
    status BOOLEAN NOT NULL,
    work_start TEXT,
    work_end TEXT
);
-- Inserção de dados iniciais
INSERT INTO users (name, profile, document, pin, coercion, card_number, status, work_start, work_end) VALUES
    ('Admin', 'admin', '12345678900', 123456, 6543210, 100001, 1, '08:00', '17:00'),
    ('João Silva', 'operator', '98765432101', 234567, 7654321, 100002, 1, '09:00', '18:00'),
    ('Maria Souza', 'operator', '11223344556', 345678, 8765432, 100003, 0, '10:00', '19:00'),
    ('Carlos Pereira', 'admin', '22334455667', 456789, 9876543, 100004, 1, '07:00', '16:00'),
    ('Ana Lima', 'operator', '33445566778', 567890, 1987654, 100005, 1, '12:00', '21:00');
