package models

type MessageChannel struct {
	messages map[int]map[int]chan string
}

var messageChannels MessageChannel

func GetMessageChannelWrapper() MessageChannel {
	if messageChannels.messages == nil {
		messageChannels.messages = make(map[int]map[int]chan string)
	}

	return messageChannels
}

/**
* Get channel by user's ID
 */
func (mc MessageChannel) GetChannel(idChatRoom int, idUser int, needCreation bool) chan string {

	if needCreation {
		if mc.messages[idChatRoom] == nil {
			mc.messages[idChatRoom] = make(map[int]chan string)
		}
		mc.messages[idChatRoom][idUser] = make(chan string)
	}

	return mc.messages[idChatRoom][idUser]
}

/**
* Get all channels
 */
func (mc MessageChannel) GetChannels() map[int]map[int]chan string {
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

func (mc MessageChannel) PutToChanel(id int, message string) {
	for _, channel := range mc.messages[id] {
		channel <- string(message)
	}
}
