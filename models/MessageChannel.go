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

func (mc MessageChannel) GetChannel(name int, needCreation bool) chan string {

	if needCreation && mc.messages[name] == nil {
		mc.messages[name] = make(chan string)
	}

	return mc.messages[name]
}
