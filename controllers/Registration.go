package controllers

/**
* Registration service
 */

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
	openKey := req.FormValue("openKey")

	if username == "" || password == "" || openKey == "" {
		result.Result = 400
		result.ResultMessage = "Login, Password and OpenKey are required"
	} else {
		profile := models.NewProfile(config.GetConnection())

		if profile.FindByUsername(username) == nil {
			profile.Username = username
			profile.SetPassword(password)
			if profile.Save() {
				profileKey := models.NewOpenKey(config.GetConnection())
				profileKey.IdProfile = profile.Id
				profileKey.Key = openKey
				profileKey.Save()
			}

			result.Data = profile
		} else {
			result.Result = 403
			result.ResultMessage = "Account exists"
		}
	}

	jsonResult, err := json.Marshal(result)

	if err != nil {
		w.WriteHeader(500)
	} else {
		w.Write(jsonResult)
	}
}
