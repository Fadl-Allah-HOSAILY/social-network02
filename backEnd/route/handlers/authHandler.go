package handlers

import (
	"backend/model"
	"backend/service"
	"database/sql"
	"net/http"
	"time"
)

func RegisterHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	var userData model.RegisterRequest
	var response model.APIResponse

	if r.Method != http.MethodPost {
		response.Code = http.StatusMethodNotAllowed
		response.Data = nil
		response.Success = false
		response.PopUp = "Invalid Method"

		WriteJSONResponse(w, response)

		return
	}

	file, header, err := r.FormFile("avatar")

	if err != nil && err != http.ErrMissingFile {
		response.Code = http.StatusBadRequest
		response.Data = nil
		response.Success = false
		response.PopUp = "Invalid avatar"

		WriteJSONResponse(w, response)

		return
	}

	birthdayString := r.FormValue("birthday")

	birthday, err := time.Parse("2006-01-02", birthdayString)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Data = nil
		response.Success = false
		response.PopUp = "Invalid birthday format"

		WriteJSONResponse(w, response)

		return
	}

	path, err := service.ImageTreatment(file, header)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Data = nil
		response.Success = false
		response.PopUp = "Invalid Avatar"

		WriteJSONResponse(w, response)

		return
	}

	userData = model.RegisterRequest{
		FirstName: r.FormValue("firstName"),
		LastName:  r.FormValue("lastName"),
		Nickname:  r.FormValue("nickname"),
		AboutMe:   r.FormValue("aboutMe"),
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
		Birthday:  birthday,
		Avatar:    path,
	}

	sessionID, err := service.Register(db, userData)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Data = nil
		response.Success = false
		response.PopUp = err.Error()

		WriteJSONResponse(w, response)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	response.Code = http.StatusCreated
	response.Data = nil
	response.Success = true
	response.PopUp = "Registration successful"

	WriteJSONResponse(w, response)
}
