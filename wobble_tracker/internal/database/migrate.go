package database

import (
    "database/sql"
    "embed"
    "fmt"
    "log"
    "path/filepath"
    "sort"
    "strings"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

// Migration represents a single migration
type Migration struct {
    ID       int
    Filename string
    SQL      string
}

// runMigrations executes all pending migrations
func runMigrations(db *sql.DB) error {
    // Create migrations table if it doesn't exist
    if err := createMigrationsTable(db); err != nil {
        return fmt.Errorf("failed to create migrations table: %w", err)
    }

    // Get all migration files
    migrations, err := loadMigrations()
    if err != nil {
        return fmt.Errorf("failed to load migrations: %w", err)
    }

    // Get applied migrations
    appliedMigrations, err := getAppliedMigrations(db)
    if err != nil {
        return fmt.Errorf("failed to get applied migrations: %w", err)
    }

    // Run pending migrations
    for _, migration := range migrations {
        if !appliedMigrations[migration.ID] {
            log.Printf("Running migration: %s", migration.Filename)
            
            if err := runMigration(db, migration); err != nil {
                return fmt.Errorf("failed to run migration %s: %w", migration.Filename, err)
            }
            
            if err := recordMigration(db, migration); err != nil {
                return fmt.Errorf("failed to record migration %s: %w", migration.Filename, err)
            }
            
            log.Printf("Migration completed: %s", migration.Filename)
        }
    }

    return nil
}

// createMigrationsTable creates the migrations tracking table
func createMigrationsTable(db *sql.DB) error {
    query := `
    CREATE TABLE IF NOT EXISTS migrations (
        id INTEGER PRIMARY KEY,
        filename TEXT NOT NULL UNIQUE,
        applied_at DATETIME DEFAULT CURRENT_TIMESTAMP
    )`
    
    _, err := db.Exec(query)
    return err
}

// loadMigrations loads all migration files from the embedded filesystem
func loadMigrations() ([]Migration, error) {
    entries, err := migrationFiles.ReadDir("migrations")
    if err != nil {
        return nil, err
    }

    var migrations []Migration
    
    for _, entry := range entries {
        if !strings.HasSuffix(entry.Name(), ".sql") {
            continue
        }

        // Extract migration ID from filename (e.g., "001_create_cats_table.sql" -> 1)
        parts := strings.Split(entry.Name(), "_")
        if len(parts) == 0 {
            continue
        }

        var id int
        if _, err := fmt.Sscanf(parts[0], "%d", &id); err != nil {
            log.Printf("Warning: couldn't parse migration ID from %s, skipping", entry.Name())
            continue
        }

        // Read migration content
        content, err := migrationFiles.ReadFile(filepath.Join("migrations", entry.Name()))
        if err != nil {
            return nil, fmt.Errorf("failed to read migration file %s: %w", entry.Name(), err)
        }

        migrations = append(migrations, Migration{
            ID:       id,
            Filename: entry.Name(),
            SQL:      string(content),
        })
    }

    // Sort migrations by ID
    sort.Slice(migrations, func(i, j int) bool {
        return migrations[i].ID < migrations[j].ID
    })

    return migrations, nil
}

// getAppliedMigrations returns a map of applied migration IDs
func getAppliedMigrations(db *sql.DB) (map[int]bool, error) {
    query := "SELECT id FROM migrations"
    
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    applied := make(map[int]bool)
    
    for rows.Next() {
        var id int
        if err := rows.Scan(&id); err != nil {
            return nil, err
        }
        applied[id] = true
    }

    return applied, rows.Err()
}

// runMigration executes a single migration
func runMigration(db *sql.DB, migration Migration) error {
    // Split SQL by semicolons and execute each statement
    statements := strings.Split(migration.SQL, ";")
    
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    for _, stmt := range statements {
        stmt = strings.TrimSpace(stmt)
        if stmt == "" || strings.HasPrefix(stmt, "--") {
            continue
        }

        if _, err := tx.Exec(stmt); err != nil {
            return fmt.Errorf("failed to execute statement: %s, error: %w", stmt, err)
        }
    }

    return tx.Commit()
}

// recordMigration records that a migration has been applied
func recordMigration(db *sql.DB, migration Migration) error {
    query := "INSERT INTO migrations (id, filename) VALUES (?, ?)"
    _, err := db.Exec(query, migration.ID, migration.Filename)
    return err
}