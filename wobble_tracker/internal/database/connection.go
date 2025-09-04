package database

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func Initialize() (*sql.DB, error) {
    dbPath := getDatabasePath()
    
    // Ensure the directory exists
    if dbPath != ":memory:" {
        dir := filepath.Dir(dbPath)
        if err := os.MkdirAll(dir, 0755); err != nil {
            return nil, err
        }
    }
    
    db, err := sql.Open("sqlite3", dbPath+"?_foreign_keys=on")
    if err != nil {
        return nil, err
    }
    
    // Test the connection
    if err := db.Ping(); err != nil {
        db.Close()
        return nil, err
    }
    
    // Configure connection pool
    db.SetMaxOpenConns(1) // SQLite works best with a single connection
    
    // Run migrations
    if err := runMigrations(db); err != nil {
        db.Close()
        return nil, err
    }
    
    return db, nil
}

// getDatabasePath returns the appropriate database path based on environment
func getDatabasePath() string {
    switch os.Getenv("APP_ENV") {
    case "test":
        return ":memory:"
    case "development":
        return "data/cats_dev.db"
    case "production":
        return "data/cats.db"
    default:
        return "data/cats.db"
    }
}