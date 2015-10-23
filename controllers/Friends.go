package controllers

/**
* Working with friends
 */

import (
	"chat/config"
	"chat/models"
	"chat/models/response"
	"encoding/json"
	"net/http"
)

type Friends struct {
}

func CreateFriends() *Friends {
	return new(Friends)
}

/**
* Get online friends
 */
func (f *Friends) GetOnlineUsers(w http.ResponseWriter, req *http.Request) {
	result := response.NewFriendsReponse(0, "")

	onlineKeys := models.GetMessageChannelWrapper().GetChannelKeys()

	result.Data = models.NewProfile(config.GetConnection()).GetUsersByIds(onlineKeys)

	jsonResult, err := json.Marshal(result)

	if err != nil {
		w.WriteHeader(500)
	} else {
		w.Write(jsonResult)
	}
}

/**
* Find profiles
*/
func (f *Friends) FindUsers(w http.ResponseWriter, req *http.Request) {
	result := response.NewFriendsReponse(0, "")

	searchStr := req.FormValue("searchStr")

	if searchStr == "" {
		result.Result = 400
		result.ResultMessage = "searchStr param is required"
		w.WriteHeader(400)
	} else {
		result.Data = models.NewProfile(config.GetConnection()).Find(searchStr)
	}

	jsonResult, err := json.Marshal(result)

	if err != nil {
		w.WriteHeader(500)
	} else {
		w.Write(jsonResult)
	}
}
