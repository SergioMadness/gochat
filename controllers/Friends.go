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
