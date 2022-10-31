package config

import "golang.org/x/net/websocket"

const MAX_USER_CONNECTIONS = 20

type ConnectionPool struct {
	connections map[int]map[int][]*websocket.Conn
}

var connectionPool *ConnectionPool

func GetConnectionPool() *ConnectionPool {
	if connectionPool == nil {
		connectionPool = new(ConnectionPool)
		connectionPool.connections = make(map[int]map[int][]*websocket.Conn)
	}
	return connectionPool
}

func AddConnection(idRoom int, idUser int, ws *websocket.Conn) error {
	var result error

	pool := GetConnectionPool()
	if pool.connections[idRoom] == nil {
		pool.connections[idRoom] = make(map[int][]*websocket.Conn)
	}
	//	if len(pool.connections[idRoom][idUser]) >= MAX_USER_CONNECTIONS {
	//		result = errors.New("Max " + string(MAX_USER_CONNECTIONS) + " connections")
	//	} else {
	pool.connections[idRoom][idUser] = nil
	pool.connections[idRoom][idUser] = append(pool.connections[idRoom][idUser], ws)
	//	}

	return result
}

func GetRoomConnections(idRoom int) map[int][]*websocket.Conn {
	return GetConnectionPool().connections[idRoom]
}

func GetUserConnections(idRoom int, idUser int) []*websocket.Conn {
	return GetConnectionPool().connections[idRoom][idUser]
}

func CloseConnection(idRoom int, idUser int) {
	pool := GetConnectionPool()
	if pool.connections[idRoom][idUser] != nil {
		for _, conn := range pool.connections[idRoom][idUser] {
			conn.Close()
		}
	}
}
