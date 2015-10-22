package controllers

import (
	"chat/config"
	"chat/helpers"
	"chat/models"
	"chat/models/response"
	"encoding/json"
	"net/http"
)

type Login struct {
}

func CreateLogin() *Login {
	return new(Login)
}

func (r Login) HandleRequest(w http.ResponseWriter, req *http.Request) {
	result := response.NewProfileResponse(0, "")

	login := req.FormValue("login")
	password := req.FormValue("password")

	if login == "" || password == "" {
		w.WriteHeader(400)
		result.Result = 400
		result.ResultMessage = "Login and Password are required"
	} else {
		profile := models.NewProfile(config.GetConnection()).FindByCredentials(login, helpers.GetMD5(password))

		result.Data = profile
	}

	jsonResult, err := json.Marshal(result)

	if err != nil {
		w.WriteHeader(500)
	} else {
		w.Write(jsonResult)
	}
}
