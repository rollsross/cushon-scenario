package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rodionross/cushon-scenario/storage"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		panic(err)
	}

	err = storage.Seed(db)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	port := "8080"

	fmt.Printf("server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
