package response

const RESPONSE_TYPE_MESSAGE int = 1
const RESPONSE_TYPE_ONLINE_FRIEND int = 2

type LongPollResponse struct {
	Response
	ResponseType int
}