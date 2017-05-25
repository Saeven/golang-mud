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

const STATE_CHARACTER_CREATION = 4

const MAX_PASSWORD_FAILURES = 3

type Connection struct {
	conn             net.Conn
	timeConnected    time.Time
	state            int8
	username         string
	server           *Server
	Player           *Player
	passwordFailures int
}

func (connection *Connection) Write(message string) {
	connection.conn.Write([]byte(message))
}

func (connection *Connection) sendMOTD() {
	connection.Write(connection.server.Motd)
	connection.Write("What is your name, mortal? ")
}

func (connection *Connection) sendMenu() {
	connection.Write(connection.server.Menu)
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

		// Player has just seen the MOTD, and is asked for username
		case STATE_LOGIN_USERNAME:
			connection.state = STATE_LOGIN_PASSWORD
			connection.username = message
			connection.Write("Your password? ")

		// Player is being asked to authenticate
		case STATE_LOGIN_PASSWORD:
			exists, player := connection.server.authenticatePlayer(connection.username, message);

			if exists {
				if player != nil { // auth succeeded
					connection.Player = player
					connection.server.onPlayerAuthenticated(connection)
				} else { // auth fails
					connection.Write("Sorry, that wasn't right. Try again: ")

					connection.passwordFailures++
					if connection.passwordFailures > MAX_PASSWORD_FAILURES {
						connection.Write( "Pfft...  Goodbye.")
						connection.conn.Close()
					}

				}
			} else {
				connection.state = STATE_CHARACTER_CREATION
			}


		default:
			connection.server.onMessageReceived(connection, message)
		}

	}
}
