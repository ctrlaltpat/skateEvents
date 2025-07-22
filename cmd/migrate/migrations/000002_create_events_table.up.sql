CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    owner_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    start_date DATETIME NOT NULL,
    end_date DATETIME NOT NULL,
    location TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'draft',
    FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE
);