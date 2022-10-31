package controllers

import (
	"fmt"
	"io"
	"net/http"
	"socialchat/config"
	"socialchat/models"

	"golang.org/x/net/websocket"
)

const COMMAND_NEW_USER = "new"
const COMMAND_RTC = "webrtc"
const COMMAND_DISCONNECT = "disconnect"

type MessagingP2P struct {
}

func CreateMessagingP2P() *MessagingP2P {
	result := new(MessagingP2P)
	//	result.connections = make(map[int][]*ConnectionStruct)
	return result
}

func (s *MessagingP2P) SendAll(roomId int, userId int, message *models.MessageP2PStruct) {
	users := config.GetRoomConnections(roomId)
	if users != nil {
		for uid, connList := range users {
			if uid != userId {
				for _, conn := range connList {
					websocket.JSON.Send(conn, message)
				}
			}
		}
	}
}

func (s *MessagingP2P) SendTo(roomId int, toId int, message *models.MessageP2PStruct) {
	users := config.GetRoomConnections(roomId)
	if users != nil {
		for uid, connList := range users {
			if uid == toId {
				for _, conn := range connList {
					websocket.JSON.Send(conn, message)
				}
			}
		}
	}
}

func (s *MessagingP2P) NewUser(roomId int, userId int) {
	s.SendAll(roomId, userId, &models.MessageP2PStruct{"new", userId, 0, userId})
}

func (s *MessagingP2P) UserDisconected(roomId int, userId int) {
	s.SendAll(roomId, userId, &models.MessageP2PStruct{"disconected", userId, 0, userId})
}

func (s *MessagingP2P) ProcessMessage(roomId int, userId int, msg *models.MessageP2PStruct) bool {
	result := false
	msg.Id = userId
	switch msg.Command {
	case COMMAND_DISCONNECT:
		s.SendAll(roomId, userId, msg)
		config.CloseConnection(roomId, userId)
		result = true
	default:
		if msg.To > 0 {
			s.SendTo(roomId, msg.To, msg)
		} else {
			s.SendAll(roomId, userId, msg)
		}
		result = true
	}
	return result
}

func (s *MessagingP2P) Start(url string) {
	// websocket handler
	onConnected := func(ws *websocket.Conn) {
		defer func() {
			ws.Close()
		}()

		token := ws.Request().FormValue("access-token")
		if token != "" {
			config.GetSession().SetCurrentUser(models.NewProfile(config.GetConnection()).GetByToken(token))
		}

		if config.GetSession().IsLoggedIn() {

			currentUserId := config.GetSession().GetCurrentUser().Id

			chatRoom := models.NewChatroom(config.GetConnection()).GetActiveRoom(currentUserId)
			s.NewUser(chatRoom.Id, currentUserId)
			config.AddConnection(chatRoom.Id, currentUserId, ws)

			for {
				msg := new(models.MessageP2PStruct)
				err := websocket.JSON.Receive(ws, msg)
				if err == io.EOF {
					return
				}
				s.ProcessMessage(chatRoom.Id, currentUserId, msg)
				fmt.Println(msg)
			}
		}
	}
	http.Handle(url, websocket.Handler(onConnected))
}
