package storage

import (
	"database/sql"

	"github.com/rodionross/cushon-scenario/helpers"
)

func Seed(db *sql.DB) error {
	err := helpers.ExecuteSQLFile(db, "./database/schemas.sql")
	if err != nil {
		return err
	}

	err = helpers.ExecuteSQLFile(db, "./database/seed.sql")
	if err != nil {
		return err
	}

	return nil
}
