package response

/**
* Response with messages
*/

import (
	"chat/models"
)

type MessageResponse struct {
	Response
	Data []models.Message
}

func NewMessageResponse(resultCode int, resultMessage string) *MessageResponse {
	result := new(MessageResponse)

	result.Result = resultCode
	result.ResultMessage = resultMessage

	return result
}
