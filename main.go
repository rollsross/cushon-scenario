package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	port := "8080"

	fmt.Printf("server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
