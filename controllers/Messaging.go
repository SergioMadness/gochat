package controllers

/**
* Messaging service
 */

import (
	"encoding/json"
	"io"
	"net/http"
	"socialchat/config"
	"socialchat/models"
)

type Messaging struct {
}

type Message struct {
	Message string
	IdUser  int
}

func CreateMessaging() *Messaging {
	return new(Messaging)
}

func (m *Messaging) HandleRequest(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("msg")
	currentUserId := config.GetSession().GetCurrentUser().Id

	chatRoom := models.NewChatroom(config.GetConnection()).GetActiveRoom(currentUserId)
	if chatRoom != nil {
		if message != "" {
			jsonResponse, _ := json.Marshal(Message{Message: message, IdUser: currentUserId})
			models.GetMessageChannelWrapper().PutToChanel(chatRoom.Id, string(jsonResponse))
		} else {
			io.WriteString(w, <-models.GetMessageChannelWrapper().GetChannel(chatRoom.Id, currentUserId, true))
		}
	}
}
