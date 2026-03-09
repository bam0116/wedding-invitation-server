package main

import (
	"database/sql"
	"net/http"

	"github.com/bam0116/wedding-invitation-server/env"
	"github.com/bam0116/wedding-invitation-server/httphandler"
	"github.com/bam0116/wedding-invitation-server/sqldb"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

func main() {
	db, err := sql.Open("sqlite3", "./sql.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqldb.SetDb(db)

	mux := http.NewServeMux()

	// Guestbook 핸들러 (OPTIONS 처리 포함)
	mux.HandleFunc("/api/guestbook", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		new(httphandler.GuestbookHandler).ServeHTTP(w, r)
	})

	// Attendance 핸들러 (OPTIONS 처리 포함)
	mux.HandleFunc("/api/attendance", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		new(httphandler.AttendanceHandler).ServeHTTP(w, r)
	})

	// CORS 설정
	corHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{env.AllowOrigin}, // .env에 https://bam0116.github.io
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := corHandler.Handler(mux)

	http.ListenAndServe(":8080", handler)
}