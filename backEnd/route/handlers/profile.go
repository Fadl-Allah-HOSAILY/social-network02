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

func ProfileHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var response model.APIResponse

	if r.Method != http.MethodGet {
		response.Code = http.StatusMethodNotAllowed
		response.Data = nil
		response.Success = false
		response.PopUp = "Invalid Method"

		WriteJSONResponse(w, response)

		return
	}

	strId := r.PathValue("id")
	profileID, err := strconv.Atoi(strId)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Data = nil
		response.Success = false
		response.PopUp = "Invalid id"

		WriteJSONResponse(w, response)

		return
	}

	viewerID, ok := authcontext.GetUserID(r.Context())
	if !ok {
		response.Code = http.StatusUnauthorized
		response.Data = nil
		response.Success = false
		response.PopUp = "Unauthorized"

		WriteJSONResponse(w, response)

		return
	}

	user, err := service.GetProfile(db, profileID, viewerID)

	if err != nil {
		response.Data = nil
		response.Success = false

		if errors.Is(err, service.ErrProfileNotFound) {
			response.Code = http.StatusNotFound
			response.PopUp = "Profile Not Found"
		} else if errors.Is(err, service.ErrPrivateProfile) {
			response.Code = http.StatusForbidden
			response.PopUp = "Private Profile"
		} else {
			response.Code = http.StatusInternalServerError
			response.PopUp = "Internal Server Error"
		}

		WriteJSONResponse(w, response)
		return
	}

	var profileResponse model.ProfileResponse

	profileResponse.UserID = user.UserID
	profileResponse.FirstName = user.FirstName
	profileResponse.LastName = user.LastName
	profileResponse.Nickname = user.Nickname
	profileResponse.Avatar = user.Avatar
	profileResponse.AboutMe = user.AboutMe
	profileResponse.ProfileVisibility = user.ProfileVisibility
	profileResponse.Birthday = user.Birthday

	response.Code = http.StatusOK
	response.Data = profileResponse
	response.Success = true
	response.PopUp = "Success"

	WriteJSONResponse(w, response)
}
