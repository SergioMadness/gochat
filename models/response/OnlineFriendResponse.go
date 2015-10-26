package response

import "chat/models"

type OnlineFriendResponse struct {
	LongPollResponse
	Data *models.Profile
}

func NewOnlineFriendResponse(resultCode int, resultMessage string) *OnlineFriendResponse {
	result:=new(OnlineFriendResponse)
	
	result.Result = resultCode
	result.ResultMessage = resultMessage
	result.ResponseType = RESPONSE_TYPE_ONLINE_FRIEND
	
	return result
}