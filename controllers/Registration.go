package controllers

import (
	"chat/config"
	"chat/models"
	"chat/models/response"
	"encoding/json"
	"net/http"
)

type Registration struct {
}

func CreateRegistration() *Registration {
	return new(Registration)
}

func (r *Registration) HandleRequest(w http.ResponseWriter, req *http.Request) {
	result := response.NewProfileResponse(0, "")

	username := req.FormValue("login")
	password := req.FormValue("password")

	if username == "" || password == "" {
		result.Result = 400
		result.ResultMessage = "Login and Password are required"
	} else {
		profile := models.NewProfile(config.GetConnection())

		if profile.FindByUsername(username) == nil {
			profile.Username = username
			profile.SetPassword(password)
			profile.Save()

			result.Data = profile
		} else {
			result.Result = 404
			result.ResultMessage = "User not found"
		}
	}

	jsonResult, err := json.Marshal(result)

	if err != nil {
		w.WriteHeader(500)
	} else {
		w.Write(jsonResult)
	}
}
