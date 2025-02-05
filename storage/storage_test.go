package storage

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rodionross/cushon-scenario/helpers"
)

func getRepoForTest(t *testing.T) (repo Repository, cleanup func()) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		panic(err)
	}

	err = helpers.ExecuteSQLFile(db, "../database/schemas.sql")
	if err != nil {
		panic(err)
	}

	err = helpers.ExecuteSQLFile(db, "../database/seed.sql")
	if err != nil {
		panic(err)
	}

	return New(db), func() {
		if err := db.Close(); err != nil {
			t.Fatalf("closing db: %s", err.Error())
		}
	}
}
