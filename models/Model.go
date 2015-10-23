package models

import (
	"database/sql"
)

type Model struct {
	conn *sql.DB
}

func (m *Model) GetConnection() *sql.DB {
	return m.conn
}

func (m *Model) SetConnection(conn *sql.DB) {
	m.conn = conn
}
