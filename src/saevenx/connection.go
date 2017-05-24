package saevenx

import (
	"net"
	"time"
	"bufio"
	"strings"
)

const STATE_WELCOME = 0
const STATE_LOGIN_USERNAME = 1
const STATE_LOGIN_PASSWORD = 2
const STATE_LOGIN_MENU = 3

type Connection struct {
	conn          net.Conn
	timeConnected time.Time
	state         int8
	username      string
	server        *Server
	player        *Player
}

func (connection *Connection) Write(message string) {
	connection.conn.Write([]byte(message))
}

func (connection *Connection) sendMOTD() {
	connection.Write(connection.server.Motd)
	connection.Write("What is your name, mortal?")
}

func (connection *Connection) listen() {
	reader := bufio.NewReader(connection.conn)

	connection.sendMOTD()
	connection.state = STATE_LOGIN_USERNAME

	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			connection.conn.Close()
			connection.server.onClientConnectionClosed(connection, err)
			return
		}

		message = strings.TrimSpace(message)

		switch connection.state {

		case STATE_LOGIN_USERNAME:
			connection.state = STATE_LOGIN_PASSWORD
			connection.username = message
			connection.Write("Your password?")

		case STATE_LOGIN_PASSWORD:
			connection.state = STATE_LOGIN_MENU
			connection.Write("Welcome. Death Awaits.\n")

		}

		connection.server.onMessageReceived(connection, message)

	}
}
