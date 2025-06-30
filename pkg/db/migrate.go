package db

import (
	"database/sql"
	"log"
	"pos-go/pkg/db/schemas"
)

// SchemaMigrator represents a single migration unit
type SchemaMigrator struct {
	Name  string
	Query string
}

func AutoMigrate(db *sql.DB) {
	migrations := []SchemaMigrator{
		{"users", schemas.SchemaUsers},
	}

	for _, m := range migrations {
		log.Printf("Migrating table: %s...", m.Name)
		if _, err := db.Exec(m.Query); err != nil {
			log.Fatalf("Migration failed for %s: %v", m.Name, err)
		}
		log.Printf("Migrated table %s successfully!", m.Name)
	}

	log.Println("✅ Semua tabel berhasil dibuat.")
	RunSeeders(db)
}
