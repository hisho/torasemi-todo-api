package migration

import (
	"database/sql"
	"embed"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4/source/httpfs"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

//go:embed migrations
var migrationFiles embed.FS

func Run(db *sql.DB) error {
	staticDir := http.FS(migrationFiles)
	src, err := httpfs.New(staticDir, "migrations")
	if err != nil {
		log.Print(err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("embedFs", src, "mysql", driver)
	if err != nil {
		return err
	}

	v, dirty, err := m.Version()
	if err != nil {
		if err == migrate.ErrNilVersion {
			log.Println("No migration present")
		} else {
			log.Fatal(err)
		}
	}

	if dirty {
		log.Fatalf("Current DIRTY version: %d, please fix issues\n", v)
	}

	log.Printf("Current version: %d\n", v)
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No change")
			return nil
		}
		log.Fatal(err)
	}
	v, _, err = m.Version()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("New version: %d\n", v)

	return nil
}
