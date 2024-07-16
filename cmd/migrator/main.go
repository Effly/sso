package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storagePath, migrationsPath, migrationsTable string

	flag.StringVar(&storagePath, "storage-path", "", "path to the database storage")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to the migrations")
	flag.StringVar(
		&migrationsTable,
		"migrations-table",
		"migrations",
		"path to the migrations table",
	)
	flag.Parse()

	if storagePath == "" { // TODO do validated func or use libs
		panic("storage path is required")
	}
	if migrationsPath == "" { // TODO do validated func or use libs
		panic("migrations path is required")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)
	if err != nil {
		panic(err)
	}

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")

			return
		}

		panic(err)
	}

	fmt.Println("applied migrations")
}
