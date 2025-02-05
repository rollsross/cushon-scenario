package helpers

import (
	"database/sql"
	"fmt"
	"os"
)

func ExecuteSQLFile(db *sql.DB, filepath string) error {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %v", filepath, err)
	}

	_, err = db.Exec(string(bytes))
	if err != nil {
		return fmt.Errorf("failed to execute SQL script: %v", err)
	}

	return nil
}
