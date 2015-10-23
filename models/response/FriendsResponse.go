package response

import (
	"chat/models"
)

type FriendsResponse struct {
	Response
	Data []models.Profile
}

func NewFriendsReponse(resultCode int, resultMessage string) *FriendsResponse {
	result := new(FriendsResponse)

	result.Result = resultCode
	result.ResultMessage = resultMessage

	return result
}
