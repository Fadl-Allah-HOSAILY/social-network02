package middleware

import (
	"backend/model"
	"backend/pkg/authcontext"
	"backend/route/handlers"
	"backend/service"
	"database/sql"
	"errors"
	"net/http"
)

func AuthMiddleware(db *sql.DB, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var response model.APIResponse

		cookie, err := r.Cookie("session_id")
		if err != nil {
			response.Code = http.StatusUnauthorized
			response.Data = nil
			response.Success = false
			response.PopUp = "Session Invalid"

			handlers.WriteJSONResponse(w, response)
			return
		}

		session, err := service.CheckSession(db, cookie.Value)
		if err != nil {
			if errors.Is(err, service.ErrSessionNotFound) ||
				errors.Is(err, service.ErrSessionExpired) {

				response.Code = http.StatusUnauthorized
				response.Data = nil
				response.Success = false
				response.PopUp = "Session Not Valid"

				handlers.WriteJSONResponse(w, response)
				return
			}

			response.Code = http.StatusInternalServerError
			response.Data = nil
			response.Success = false
			response.PopUp = "Internal Server Error"

			handlers.WriteJSONResponse(w, response)
			return
		}

		ctx := authcontext.WithUserID(r.Context(), session.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
