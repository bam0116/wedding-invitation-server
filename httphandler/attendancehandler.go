package httphandler

import (
	"encoding/json"
	"net/http"

	"github.com/bam0116/wedding-invitation-server/env"
	"github.com/bam0116/wedding-invitation-server/sqldb"
	"github.com/bam0116/wedding-invitation-server/types"
)

type AttendanceHandler struct {
	http.Handler
}

func (h *AttendanceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// CORS 헤더
	w.Header().Set("Access-Control-Allow-Origin", env.AllowOrigin)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		var attendance types.AttendanceCreate
		if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := sqldb.CreateAttendance(attendance.Side, attendance.Name, attendance.Meal, attendance.Count); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	if r.Method == http.MethodGet {
    rows, err := sqldb.GetDb().Query("SELECT id, side, name, meal, count, created_at FROM attendance ORDER BY created_at DESC")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var attendances []types.Attendance
    for rows.Next() {
        var a types.Attendance
        if err := rows.Scan(&a.ID, &a.Side, &a.Name, &a.Meal, &a.Count, &a.CreatedAt); err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        attendances = append(attendances, a)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(attendances)
    return
	}
}