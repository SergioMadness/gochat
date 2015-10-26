package models

type MessageChannel struct {
	messages map[int]chan string
}

var messageChannels MessageChannel

func GetMessageChannelWrapper() MessageChannel {
	if messageChannels.messages == nil {
		messageChannels.messages = make(map[int]chan string)
	}

	return messageChannels
}

/**
* Get channel by user's ID
*/
func (mc MessageChannel) GetChannel(name int, needCreation bool) chan string {

	if needCreation && mc.messages[name] == nil {
		mc.messages[name] = make(chan string)
	}

	return mc.messages[name]
}

/**
* Get all channels
*/
func (mc MessageChannel) GetChannels() map[int]chan string {
	return mc.messages
}

/**
* Get online users
*/
func (mc MessageChannel) GetChannelKeys() []int {
	var result []int

	for key := range messageChannels.messages {
		result = append(result, key)
	}

	return result
}
