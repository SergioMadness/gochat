package response

import (
	"chat/models"
)

type ProfileResponse struct {
	Response

	Data *models.Profile
}

func NewProfileResponse(resultCode int, resultMessage string) *ProfileResponse {
	result := new(ProfileResponse)

	result.Result = resultCode
	resultMessage = resultMessage

	return result
}
