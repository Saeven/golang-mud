package saevenx

import (
	"net"
	"time"
	"bufio"
)

type Connection struct {
	conn net.Conn
	timeConnected time.Time
	state int8
	server *Server
	player *Player
}

func (connection *Connection) Write(message string) {
	connection.conn.Write([]byte(message))
}

func (connection *Connection) listen() {
	reader := bufio.NewReader(connection.conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			connection.conn.Close()
			connection.server.onClientConnectionClosed(connection, err)
			return
		}
		connection.server.onMessageReceived(connection, message)

	}
}

