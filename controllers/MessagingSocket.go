package controllers

import (
	"io"
	"net/http"
	"socialchat/config"
	"socialchat/models"

	"golang.org/x/net/websocket"
)

type MessagingSocket struct {
	connections map[int][]*websocket.Conn
}

func CreateMessagingSocket() *MessagingSocket {
	result := new(MessagingSocket)
	result.connections = make(map[int][]*websocket.Conn)
	return result
}

func (s *MessagingSocket) AddConnection(roomId int, ws *websocket.Conn) {
	if s.connections[roomId] == nil {
		s.connections[roomId] = []*websocket.Conn{}
	}
	s.connections[roomId] = append(s.connections[roomId], ws)
}

func (s *MessagingSocket) SendAll(roomId int, message models.MessageStruct) {
	if s.connections[roomId] != nil {
		for _, conn := range s.connections[roomId] {
			websocket.JSON.Send(conn, message)
		}
	}
}

func (s *MessagingSocket) Start(url string) {
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
			s.AddConnection(chatRoom.Id, ws)

			for {
				var msg models.MessageStruct
				err := websocket.JSON.Receive(ws, &msg)
				if err == io.EOF {
					return
				}
				s.SendAll(chatRoom.Id, msg)
			}
		}
	}
	http.Handle(url, websocket.Handler(onConnected))
}
