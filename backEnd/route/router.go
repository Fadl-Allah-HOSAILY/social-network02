package route

import (
	"backend/route/handlers"
	"backend/route/middleware"
	"database/sql"
	"net/http"
)

func NewRouter(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterHandler(db, w, r)
	})
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(db, w, r)
	})
	mux.Handle("GET /api/profile/{id}",
		middleware.AuthMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handlers.ProfileHandler(db, w, r)
		})),
	)
	mux.Handle("POST /api/follow/{id}",
		middleware.AuthMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handlers.FollowHandler(db, w, r)
		})),
	)
	mux.Handle("GET /api/followers/{id}",
		middleware.AuthMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handlers.FollowFetchHandler(db, w, r, "followers")
		})),
	)
	mux.Handle("GET /api/following/{id}",
		middleware.AuthMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handlers.FollowFetchHandler(db, w, r, "following")
		})),
	)

	return mux
}
