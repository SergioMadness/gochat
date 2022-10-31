package response

/**
* Response with messages
*/

import (
	"socialchat/models"
)

type MessageResponse struct {
	LongPollResponse
	Data []models.Message
}

func NewMessageResponse(resultCode int, resultMessage string) *MessageResponse {
	result := new(MessageResponse)

	result.ResponseType = RESPONSE_TYPE_MESSAGE
	result.Result = resultCode
	result.ResultMessage = resultMessage

	return result
}
