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
	mux.Handle("/api/guestbook", new(httphandler.GuestbookHandler))
	mux.Handle("/api/attendance", new(httphandler.AttendanceHandler))

	corHandler := cors.New(cors.Options{
    AllowedOrigins:   []string{env.AllowOrigin},
    AllowedMethods:   []string{
        http.MethodGet,
        http.MethodPost,
        http.MethodPut,
        http.MethodOptions, // ← 이거 추가
    },
    AllowedHeaders:   []string{"*"}, // 모든 헤더 허용
    AllowCredentials: true,
})
	handler := corHandler.Handler(mux)

	http.ListenAndServe(":8080", handler)
}
