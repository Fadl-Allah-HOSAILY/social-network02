package handlers

import (
	"backend/model"
	"backend/pkg/authcontext"
	"backend/service"
	"database/sql"
	"errors"
	"net/http"
	"strconv"
)

func FollowHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var response model.APIResponse

	if r.Method != http.MethodPost {
		response.Code = http.StatusMethodNotAllowed
		response.Data = nil
		response.Success = false
		response.PopUp = "Invalid Method"

		WriteJSONResponse(w, response)

		return
	}

	strId := r.PathValue("id")
	followingID, err := strconv.Atoi(strId)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Data = nil
		response.Success = false
		response.PopUp = "Invalid id"

		WriteJSONResponse(w, response)

		return
	}

	if followingID <= 0 {
		response.Code = http.StatusBadRequest
		response.Data = nil
		response.Success = false
		response.PopUp = "Bad Request"

		WriteJSONResponse(w, response)

		return
	}

	followerID, ok := authcontext.GetUserID(r.Context())
	if !ok {
		response.Code = http.StatusUnauthorized
		response.Data = nil
		response.Success = false
		response.PopUp = "Unauthorized"

		WriteJSONResponse(w, response)

		return
	}

	action, err := service.FollowUser(db, followerID, followingID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrCannotFollowSelf):
			response.Code = http.StatusBadRequest
			response.PopUp = "Cannot follow yourself"

		case errors.Is(err, service.ErrFollowerNotFound),
			errors.Is(err, service.ErrFollowingNotFound):
			response.Code = http.StatusNotFound
			response.PopUp = "User not found"

		default:
			response.Code = http.StatusInternalServerError
			response.PopUp = "Internal Server Error"
		}

		response.Data = nil
		response.Success = false

		WriteJSONResponse(w, response)
		return
	}

	response.Code = http.StatusOK
	response.Data = nil
	response.Success = true

	if action == "followed" {
		response.PopUp = "Followed successfully"
	} else {
		response.PopUp = "Unfollowed successfully"
	}

	WriteJSONResponse(w, response)
}
