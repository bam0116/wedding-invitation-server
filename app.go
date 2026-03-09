package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/bam0116/wedding-invitation-server/env"
	"github.com/bam0116/wedding-invitation-server/httphandler"
	"github.com/bam0116/wedding-invitation-server/sqldb"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

func main() {
	db, err := sql.Open("sqlite3", "./sql.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqldb.SetDb(db)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/guestbook", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		new(httphandler.GuestbookHandler).ServeHTTP(w, r)
	})

	mux.HandleFunc("/api/attendance", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		new(httphandler.AttendanceHandler).ServeHTTP(w, r)
	})

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{env.AllowOrigin},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}