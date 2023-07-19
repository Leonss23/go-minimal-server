package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Multiplexes requests onto a single handler.
	m := http.NewServeMux()

	const addr = ":8000"

	m.HandleFunc("/", handlePage)

	// http://localhost:8000/
	srv := http.Server{
		Addr:         addr,
		Handler:      m,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("Server running on port", addr)

	err := srv.ListenAndServe()

	log.Fatalln(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request\nURL: %v\nMethod: %v\nBody: %v\n\n", r.URL, r.Method, r.Body)
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	const page = `<html>
	<body>
	<main>
	<h1>
	Hello from my Go server!
	</h1>
	</main>
	</body>
	</html>`
	w.WriteHeader(200)
	bytePage := []byte(page)
	w.Write(bytePage)
}
