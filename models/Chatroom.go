package models

import (
	"database/sql"
	"fmt"
)

type Chatroom struct {
	Model
	Id       int
	IsActive bool
	ClosedAt int64
}

func NewChatroom(conn *sql.DB) *Chatroom {
	result := new(Chatroom)
	result.SetConnection(conn)
	return result
}

func (c *Chatroom) GetActiveRoom(idUser int) *Chatroom {
	err := c.GetConnection().QueryRow("SELECT id, is_active, closed_at FROM chatroom WHERE is_active = true AND id IN (SELECT id_chatroom FROM chatusers WHERE id_user = $1)", idUser).Scan(&c.Id, &c.IsActive, &c.ClosedAt)
	if err == nil {
		return nil
	}
	return c
}

func (c *Chatroom) Save() *Chatroom {
	stmt, err := c.GetConnection().Prepare("INSERT INTO chatroom (`created_at`, `is_active`) VALUES (LOCALTIME, 1)")

	if err != nil {
		fmt.Println(err)
	} else {
		_, err = stmt.Exec()
	}

	return c
}
