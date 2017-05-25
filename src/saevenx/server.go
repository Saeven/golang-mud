package saevenx

import (
	"net"
	"time"
	"fmt"
	"io/ioutil"
	"os"
)

type Server struct {
	connectionList []*Connection
	playerList     []*Connection
	ticker         *time.Ticker
	Motd           string
	Menu           string
}

func CreateServer() *Server {
	server := &Server{
		connectionList: make([]*Connection, 0),
		playerList:     make([]*Connection, 0),
	}

	pwd, _ := os.Getwd()

	fmt.Printf("[CONFIG] Current working directory set to %s\n", pwd)

	// 1. Pull in the welcome screen
	fmt.Println("[CONFIG] Pulling MOTD")
	motdBytes, _ := ioutil.ReadFile(pwd + "/resources/MOTD")
	server.Motd = string(motdBytes)

	// 2. Pull in the menu
	fmt.Println("[CONFIG] Pulling Menu")
	menuBytes, _ := ioutil.ReadFile(pwd + "/resources/Menu")
	server.Menu = string(menuBytes)

	return server
}

func (server *Server) AddConnection(connection net.Conn) *Connection {
	newConnection := Connection{conn: connection, timeConnected: time.Now(), server: server, state: STATE_WELCOME}
	server.connectionList = append(server.connectionList, &newConnection)
	go newConnection.listen()
	fmt.Printf("[CONN] There are %d connected users.\n", server.ConnectionCount())
	return &newConnection
}

func (server *Server) onClientConnectionClosed(connection *Connection, err error) {
	// delete the connection

	for i, conn := range server.connectionList {
		if conn == connection {
			server.connectionList[i] = server.connectionList[len(server.connectionList)-1]
			server.connectionList[len(server.connectionList)-1] = nil
			server.connectionList = server.connectionList[:len(server.connectionList)-1]
			break
		}
	}

	fmt.Printf("[DISC] There are %d connected users.\n", server.ConnectionCount())
}

func (server *Server) authenticatePlayer(username string, password string) (bool, *Player) {

	if !userExists(username) {
		return false, nil
	}

	player := authenticate(username, password)

	if player != nil {
		return true, player
	}

	return true, nil
}

/**
 * Check the database to see if the player exists
 */
func userExists(username string) bool {
	return username == "Saeven"
}

func authenticate(username string, password string) *Player {
	if username == "Saeven" && password == "123" {
		return &Player{Name: username}
	}
	return nil
}

func (server *Server) onPlayerAuthenticated(connection *Connection) {
	fmt.Printf("[AUTH] Player authenticated (%s)\n", connection.Player.Name)

	connection.state = STATE_LOGIN_MENU
	connection.Write("Welcome. Death Awaits.\n")
	connection.sendMenu()
}

/**
 * Start the main game loop
 */
func (server *Server) Start() {
	server.ticker = time.NewTicker(time.Millisecond * 3000)

	go func() {
		for range server.ticker.C {
			for _, c := range server.connectionList {
				fmt.Printf("[TICK] Running update tick on player (%s) at state [%d]\n", c.username, c.state)
			}
		}
	}()

}

func (server *Server) ConnectionCount() int {
	return len(server.connectionList)
}

/**
 * A message has been received on a given connection descriptor
 */
func (server *Server) onMessageReceived(connection *Connection, message string) {
	connection.Write("I received this string :" + message)
}
