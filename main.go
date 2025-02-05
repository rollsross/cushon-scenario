package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	port := "8080"

	fmt.Printf("server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
