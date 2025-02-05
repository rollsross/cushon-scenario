package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rodionross/cushon-scenario/helpers"
	"github.com/rodionross/cushon-scenario/server"
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

	err = helpers.ExecuteSQLFile(db, "database/schemas.sql")
	if err != nil {
		panic(err)
	}

	err = helpers.ExecuteSQLFile(db, "database/seed.sql")
	if err != nil {
		panic(err)
	}

	store := storage.New(db)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/isa-account/{id}/create", server.HandleCreateAccountAndFund(store))
	mux.HandleFunc("GET /api/isa-account/{id}", server.HandleGetAccountAndFund(store))

	port := "8080"

	fmt.Printf("server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
