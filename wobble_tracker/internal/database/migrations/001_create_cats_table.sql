CREATE TABLE IF NOT EXISTS cats (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    breed TEXT,
    color TEXT,
    age INTEGER,
    weight REAL,
    is_indoor BOOLEAN DEFAULT 1,
    is_vaccinated BOOLEAN DEFAULT 0,
    microchip_id TEXT UNIQUE,
    owner_name TEXT,
    owner_email TEXT,
    notes TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for common queries
CREATE INDEX IF NOT EXISTS idx_cats_name ON cats(name);
CREATE INDEX IF NOT EXISTS idx_cats_breed ON cats(breed);
CREATE INDEX IF NOT EXISTS idx_cats_owner_email ON cats(owner_email);
CREATE INDEX IF NOT EXISTS idx_cats_microchip ON cats(microchip_id);

-- Trigger to automatically update updated_at timestamp
CREATE TRIGGER IF NOT EXISTS update_cats_updated_at 
    AFTER UPDATE ON cats
    FOR EACH ROW
BEGIN
    UPDATE cats SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;