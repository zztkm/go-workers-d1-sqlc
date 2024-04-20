DROP TABLE IF EXISTS todos;
CREATE TABLE IF NOT EXISTS todos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT current_timestamp -- utc
);
CREATE INDEX idx_todos_on_created_at ON todos (created_at DESC);
