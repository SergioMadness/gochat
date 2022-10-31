package models

/**
* Message model
 */

import (
	"database/sql"
	"fmt"
	"log"
)

type Message struct {
	Model

	id         int
	IdChatroom int
	IdUser     int
	Message    string
	CreateDate int64
	IsRead     bool
}

func NewMessage(conn *sql.DB) *Message {
	result := new(Message)

	result.SetConnection(conn)

	return result
}

func (m *Message) GetUndeliveredMessages(idChatroom int) []Message {
	var result []Message

	rows, err := m.GetConnection().Query("SELECT * FROM chatmessage WHERE is_delivered=0 AND id_chatroom=?", idChatroom)

	if err == nil {
		for rows.Next() {
			var message Message
			if err := rows.Scan(&message.id, &message.IdUser, &message.IdChatroom, &message.Message, &message.CreateDate, &message.IsRead); err != nil {
				log.Fatal(err)
			} else {
				result = append(result, message)
			}
		}
	} else {
		log.Fatal(err)
	}

	return result
}

func (m *Message) Save() bool {
	stmt, err := m.GetConnection().Prepare("INSERT INTO chatmessage (`id_user`, `id_chatroom`, `message`, `created_at`, `is_delivered`) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		fmt.Println(err)
	} else {
		_, err = stmt.Exec(m.IdUser, m.IdChatroom, m.Message, m.CreateDate, m.IsRead)
	}

	if err == nil {
		return false
	}

	return true
}
