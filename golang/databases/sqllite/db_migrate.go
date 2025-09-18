package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	"github.com/golang-migrate/migrate/source/file"
)



func db_migrate(db *sql.DB, dir string) {

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		panic(err)
	}

	fScr, err := (&file.File{}).Open("./migrations")
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithInstance("file", fScr, "sqlite3", instance)
	if err != nil {
		panic(err)
	}
	switch dir {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	default:
		panic("unknown command: " + dir)
	}
	fmt.Print(err)
}