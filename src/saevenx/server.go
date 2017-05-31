package saevenx

import (
	"net"
	"time"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Server struct {
	connectionList []*Connection
	playerList     []*Connection
	roomList       map[int]*Room
	ticker         *time.Ticker
	Motd           string
	Menu           string
}

var ServerInstance *Server

func GetServer() *Server {

	if ServerInstance == nil {
		ServerInstance = &Server{
			connectionList: make([]*Connection, 0),
			playerList:     make([]*Connection, 0),
		}

		pwd, _ := os.Getwd()

		fmt.Printf("[CONFIG] Current working directory set to %s\n", pwd)

		// 1. Pull in the welcome screen
		fmt.Println("[CONFIG] Pulling MOTD")
		motdBytes, _ := ioutil.ReadFile(pwd + "/resources/MOTD")
		ServerInstance.Motd = string(motdBytes)

		// 2. Pull in the menu
		fmt.Println("[CONFIG] Pulling Menu")
		menuBytes, _ := ioutil.ReadFile(pwd + "/resources/Menu")
		ServerInstance.Menu = string(menuBytes)

		// 3. Prepare the command hashes
		fmt.Println("[CONFIG] Preparing commands")
		prepareCommands()

		// 4. Load in the rooms
		fmt.Println("[CONFIG] Loading rooms")
		ServerInstance.roomList = loadRooms()
	}

	return ServerInstance
}

func (server *Server) AddConnection(connection net.Conn) *Connection {
	newConnection := Connection{conn: connection, timeConnected: time.Now(), state: STATE_WELCOME}
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

/**
 * User log-in process
 */
func (server *Server) login(username string, password string) (bool, *Player) {

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
 * @TODO perform an actual database scan
 */
func userExists(username string) bool {
	return username == "Saeven"
}

/**
 * Authenticate a user via database, and fetch the player
 * @TODO implement the actual check!
 */
func authenticate(username string, password string) *Player {
	if username == "Saeven" && password == "123" {
		player := &Player{Name: username, CurrentRoom: 1, hitPointsMax: 100, hitPoints: 50, vitalityMax: 250, vitality: 250, race: getRace("demon")}
		player.inventory = []*Item{
			{Name: "A Dark Sword", Description: "A test object to test object loading"},
		}

		return player

	}
	return nil
}

/**
 * Server-side trigger when authentication occurs in the comm handler
 */
func (server *Server) onPlayerAuthenticated(connection *Connection) {
	fmt.Printf("[AUTH] Player authenticated (%s)\n", connection.Player.Name)
	server.playerList = append(server.playerList, connection)

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
			for _, c := range server.playerList {
				fmt.Printf("[TICK] Running update tick on player (%s) at state [%d]\n", c.username, c.state)
				if c.Player != nil {
					c.Player.pulseUpdate()
				}
			}
		}
	}()

}

/**
 * How many connections are active?
 */
func (server *Server) ConnectionCount() int {
	return len(server.connectionList)
}

/**
 * A message has been received on a given connection descriptor
 */
func (server *Server) onMessageReceived(connection *Connection, message string) {

	if len(message) == 0 {
		connection.Player.sendPrompt()
		return
	}
	words := strings.Fields(message)
	input, arguments := words[0], words[1:]

	connection.Player.do(input, arguments)
	connection.Player.sendPrompt()
}

func (server *Server) getRoom(roomId int) *Room {
	return server.roomList[roomId]
}
