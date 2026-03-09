package httphandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bam0116/wedding-invitation-server/env"
	"github.com/bam0116/wedding-invitation-server/sqldb"
	"github.com/bam0116/wedding-invitation-server/types"
)

type GuestbookHandler struct {
	http.Handler
}

func (h *GuestbookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 모든 요청에 CORS 헤더 추가
	w.Header().Set("Access-Control-Allow-Origin", env.AllowOrigin)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case http.MethodGet:
		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		guestbook, err := sqldb.GetGuestbook(offset, limit)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(guestbook)

	case http.MethodPost:
		var post types.GuestbookPostForCreate
		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := sqldb.CreateGuestbookPost(post.Name, post.Content, post.Password); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	case http.MethodPut:
		var post types.GuestbookPostForDelete
		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := sqldb.DeleteGuestbookPost(post.Id, post.Password); err != nil {
			if err.Error() == "INCORRECT_PASSWORD" {
				w.WriteHeader(http.StatusForbidden)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}